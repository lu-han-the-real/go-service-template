# go-service-template
Service Template of Go Microservices

## About This Repo
This is an example microservice providing Get/Post API endpoint using native [go-kit](https://gokit.io/) library.

## Environment Setup
1. Ensure you've installed golang following [installation tutorial](https://jimkang.medium.com/install-go-on-mac-with-homebrew-5fa421fc55f5)
2. download this repo using git.
3. under the source code directory (e.g.: $HOME/Code/go/src/go-service-template
   ), run `go install ./...` to compile it

## Start Server and Send Requests
1. run `go-service-template` to start the server, confirm the server runs at: `localhost:8080` as default address
2. choose your favorite api tool to send sample request. 
   1. [curl](https://curl.se/docs/httpscripting.html)
   2. [postman](https://www.postman.com/)