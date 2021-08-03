# Questions about Golang

1. Why does `go build main.go` and `go build *.go` output a .exe file (Windows specific), but then I run `go run main.go` and it doesn't seem to be used?

2. Are go.mod and go.sum essentially the JS equivalent of a package.json file?

3. I had some trouble trying to move my Person struct to another file. How can I move it to a person.go and have the app still run? Packages and imports did not work how I expected them to.

4. The internet is suggesting I use a package called mux by gorilla for HTTP routing. Is this package commonly used by businesses for production code? If so, what are the shortcomings in Go's built in HTTP handling that it was built to address?