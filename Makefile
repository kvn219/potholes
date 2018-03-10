all: compile docker push clean

compile:
	@echo compile

docker_build:
	@echo building docker image
	docker build --force-rm=true -t kvn219/potholes .

docker_run:
	@echo run docker container
	docker run --name potholes --rm -v $(pwd):/go/src/potholes/outputs -it -d kvn219/potholes
	docker exec -it potholes sh

push:
	@echo push

clean:
	@echo clean
	docker rmi $(docker images -q --filter "dangling=true") -f

video:
	@echo make video!
	asciinema rec

gif:
	docker run --rm -v $(PWD):/data asciinema/asciicast2gif -t asciinema -s 1 -S 1 -w 100 -h 50 167872.cast demo.gif


test:
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out
