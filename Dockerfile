FROM golang:1.16

WORKDIR /go/src/app
COPY . .

RUN apt-get update \
    && apt install -y protobuf-compiler \
    && go get google.golang.org/protobuf/cmd/protoc-gen-go \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go
