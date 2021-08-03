# Go Rest API

Investigating the Go programming language.

## Running locally without Docker

Run `go main.go` in the `src` folder.

## Running locally with Docker

Run `docker build --tag go-rest-api .` to create the image.

Run `docker images` to verify the image exists.

Run `docker run -p 10000:10000 go-rest-api` to start the container.

Run `docker ps` to verify the container exists.

Run `docker inspect [container Id]` for additional information on the container.