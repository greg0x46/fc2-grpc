current_dir := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

build:
	docker build -t grpc-example .
up:
	docker run -d --name grpc-example --rm -v "$(current_dir)/.:/go/src/app" grpc-example
bash: 
	docker run -it --name grpc-example --rm -v "$(current_dir)/.:/go/src/app" grpc-example bash
