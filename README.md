# Go Rest API

Investigating the Go programming language.

## Running Locally

Create a `.env` file in the root of this project. Copy the contents of `.env.token` into it, and fill in the values.

### With Docker

Run `docker build --tag go-rest-api .` to create the image.

Run `docker images` to verify the image exists.

Run `docker run -p 10000:10000 go-rest-api` to start the container.

Run `docker ps` to verify the container exists.

Run `docker inspect [container Id]` for additional information on the container.