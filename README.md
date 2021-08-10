# Go Rest API

Investigating the Go programming language. [Victor Steven's 'Go-JWT-Postgres-Mysql-Restful-API'](https://github.com/victorsteven/Go-JWT-Postgres-Mysql-Restful-API) was a great resource and reference for this project.

## Running Locally

Create a `.env` file in the root of this project. Copy the contents of `.env.token` into it, and fill in the values.

### With Docker

Run `docker-compose up -d` to create and start all necessary containers and services.

I recommend running `go build` and fixing all errors before attempting the docker command.