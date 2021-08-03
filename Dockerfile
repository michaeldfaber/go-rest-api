FROM golang:1.16.6-alpine3.14
WORKDIR /app
COPY src/go.mod ./
COPY src/go.sum ./
RUN go mod download
COPY src/*.go ./
RUN go build -o /go-rest-api
EXPOSE 10000
CMD [ "/go-rest-api" ]