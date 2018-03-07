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
	docker run --rm -v $(PWD):/data asciinema/asciicast2gif -t asciinema -s 1 -S 2 -w 200 -h 40 potholes/demo.cast demo.gif
