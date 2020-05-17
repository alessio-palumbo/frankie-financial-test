# Frankie Financial - Coding Test

## Project description

Build a simple RESTful service based of a Swagger/OpenAPI v2.0 definition.<br>
There is a single POST endpoint that takes a JSON payload which needs to be validated
according to predefined rules.<br>
If the validation is successful then a simple JSON structure is returned, otherwise
an error response returns a description of all issues found.

## Setup

### Installing Go (1.12 or higher)

Install Go following the official instructions: https://golang.org/doc/install

### Using Makefile

#### Local build

* Run `make run-build` to compile and run the generated executable
* Run `make run` to build and run

* Run `make test` to run all the test
* Run `make test-cover` to run print out the test coverage

* Run `make clean` to clean up

#### Docker build

* Run `make docker-build` to build the image from the Dockerfile
* Run `make docker-run` to run the image in a _detached_ container listening on port 8080
  * Example curl for testing:<br>
    `curl -X POST localhost:8080/isgood -d '[{"checkType":"DEVICE","activityType":"LOGIN","checkSessionKey":"10001","activityData":[{"kvpKey":"ip.address","kvpValue":"1.23.45.123","kvpType":"general.string"}]}]'`
* Run `make docker-clean` to stop and remove container and image

### Without Makefile

#### Local build

* Run `go build -o main` to build and then run executable with `./main`
* Run `go run main.go` to build and run

* Run `go test -v ./...` to run all the tests
* To check test coverage run:
  * `go test -coverprofile cover.out ./... && go tool cover -func cover.out && rm cover.out`

* Run `go clean` to clean up

#### Docker build

* Run `docker build -t frankie-image -f ./images/Dockerfile .` to build image.
* Run `docker container run -d --rm -p 8080:8080 --name frankie frankie-image` to run the container
* To clean up run:
  * docker ps -aq -f name=frankie | xargs docker stop || true; docker image rm
  * go clean

## Swagger docs

### With Makefile

* Run `make swagger-setup` to install swaggo and re-generate docs
* Run `make swagger` to re-generate docs after a change

### Without Makefile

* Run `go get -u github.com/swaggo/swag/cmd/swag` to install swaggo
* Run `swag init` to generate swagger docs

### To check docs

* Terminal: `curl localhost:8080/swagger/doc.json`
* Browser: `http://localhost:8080/swagger/index.html`
* Files: `docs/swagger.json`, `docs/swagger.yaml`
