FROM golang:1.16.6-alpine3.14 as builder
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-rest-api .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY --from=builder /app/go-rest-api .
COPY --from=builder /app/.env .
EXPOSE 10000
CMD [ "./go-rest-api" ]