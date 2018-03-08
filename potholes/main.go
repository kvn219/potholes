// This program downloads NYC 311 pothole complaints since 2010.
// More information about the data can be found here:
// https://nycopendata.socrata.com/Social-Services/311-Service-Requests-from-2010-to-Present/erm2-nwe9
package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// Complaint data structure.
type Complaint struct {
	CreatedDate           string `json:"created_date"`
	Agency                string `json:"agency"`
	AgencyName            string `json:"agency_name"`
	ComplaintType         string `json:"complaint_type"`
	Description           string `json:"descriptor"`
	UniqueKey             string `json:"unique_key"`
	CommunityBoard        string `json:"community_board"`
	ResolutionDescription string `json:"resolution_description"`
	StreetName            string `json:"street_name"`
	Latitude              string `json:"latitude"`
	Longitude             string `json:"longitude"`
}

func main() {
	fmt.Println("Getting the latest NYC pothole complaints...")
	resp := getReqFromSocrata()
	// Read response from Socrata.
	body := readResp(resp)
	// Initialize complaint data type.
	complaint := []Complaint{}
	// Unmarshall request.
	err := json.Unmarshal(body, &complaint)
	checkError("Cannot Unmarshall json!\n", err)
	// Create output file.
	file := createFile("results.csv")
	// Wait until file creation is finished.
	defer file.Close()
	// Create a csv writer.
	writer := csv.NewWriter(file)
	// Wait till the writer creation is complete.
	defer writer.Flush()
	// Create headers for the csv file.
	headers := []string{
		"record",
		"created_date",
		"unique_key",
		"agency_name",
		"description",
		"street_name",
		"resolution_description",
		"community_board",
		"latitude",
		"longitude",
	}
	// Initialize the header data type.
	var header []string
	// Append the header to csv file.
	for _, v := range headers {
		header = append(header, v)
	}
	// Write the header to the csv file.
	writer.Write(header)
	// Loop through each complaint and grab key information.
	for idx := range complaint {
		// Initialize record data type.
		var record []string
		// Loop index
		fmt.Println("--------------------")
		i := strconv.Itoa(idx)
		fmt.Println("Complaint: ", idx)
		record = append(record, i)
		fmt.Println("--------------------")
		// Unique Key
		fmt.Println("\tUnique Key: \t\t  ", complaint[idx].UniqueKey)
		record = append(record, complaint[idx].UniqueKey)
		// Created Date
		fmt.Println("\tCreated Date: \t\t  ", complaint[idx].CreatedDate)
		record = append(record, complaint[idx].CreatedDate)
		// Agency
		record = append(record, complaint[idx].Agency)
		// Agency Name
		fmt.Println("\tAgency Name: \t\t  ", complaint[idx].AgencyName)
		record = append(record, complaint[idx].AgencyName)
		// Description
		record = append(record, complaint[idx].Description)
		// Street Name
		fmt.Println("\tStreet Name: \t\t  ", complaint[idx].StreetName)
		streetName := strings.TrimSpace(complaint[idx].StreetName)
		record = append(record, streetName)
		// Resolution Description
		fmt.Println("\tResolution Description:")
		fmt.Printf("\t%s\n", complaint[idx].ResolutionDescription)
		record = append(record, complaint[idx].ResolutionDescription)
		// Add the rest of information we want.
		record = append(record, complaint[idx].CommunityBoard)
		record = append(record, complaint[idx].Latitude)
		record = append(record, complaint[idx].Longitude)
		// Write row to csv file.
		err := writer.Write(record)
		checkError("Cannot write to file", err)
	}
}

// userArgs grabs user arguments from the command line.
func userArgs() string {
	// Limit of records.
	l := flag.Int("limit", 10, "Number of records to pull.")
	// Order inwhich to sort the records.
	o := flag.String("order", "DESC", "Sort records in {DESC|ASC} order.")
	flag.Parse()
	// Construct the query string.
	baseURL := "https://data.cityofnewyork.us/resource/fhrw-4uyv.json"
	complainType := "?$where=descriptor%20=%20%27Pothole%27"
	limit := fmt.Sprintf("&$limit=%d", *l)
	orderBy := fmt.Sprintf("&$order=created_date %s", *o)
	orderByFmt := strings.Replace(orderBy, " ", "%20", -1)
	// Combine the base url and query strings
	url := baseURL + complainType + limit + orderByFmt
	return url
}

// checkError checks and log errors.
func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

// getReqFromSocrata grabs user arguments and makes a request.
func getReqFromSocrata() *http.Response {
	url := userArgs()
	resp, err := http.Get(url)
	checkError("Cannot get request from Sorcrata!", err)
	return resp
}

// readResp reads the response body and checks for errors.
func readResp(resp *http.Response) []byte {
	body, err := ioutil.ReadAll(resp.Body)
	checkError("Cannot read response body.", err)
	defer resp.Body.Close()
	return body
}

// createFile creates a file and checks for errors.
func createFile(fname string) *os.File {
	file, err := os.Create(fname)
	checkError("Cannot create file", err)
	return file
}
