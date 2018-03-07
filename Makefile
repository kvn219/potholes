all: compile docker push clean

compile:
	@echo compile

docker:
	@echo docker
	docker build --force-rm=true -t kvn219/potholes:single .

push:
	@echo push

clean:
	@echo clean

video:
	@echo make video!
	asciinema rec

gif:
	docker run --rm -v $(PWD):/data asciinema/asciicast2gif -t asciinema -s 1 -S 1 -w 100 -h 50 167872.cast demo.gif
