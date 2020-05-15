APP_NAME   = frankie-financial-test
DOCKERFILE = images/Dockerfile
COMMIT_SHA = $(shell git rev-parse --short HEAD)
IMAGE      = $(APP_NAME):$(COMMIT_SHA)

.PHONY: build
build: clean
	go build -o $(APP_NAME) main.go

.PHONY: run-build
run-build: build
	./$(APP_NAME)

.PHONY: run
run:
	go run main.go

.PHONY: test
test:
	go test -v ./...

.PHONY: test-cover
test-cover:
	go test -coverprofile cover.out ./... && \
		go tool cover -func cover.out && \
		rm cover.out

.PHONY: clean
clean:
	go clean

# Docker

.PHONY: docker-build
docker-build: build
	docker build -t $(IMAGE) -f $(DOCKERFILE) .

.PHONY: docker-run
docker-run: docker-build
	docker container run -d --rm -p 8080:8080 --name $(APP_NAME) $(IMAGE)

.PHONY: docker-clean
docker-clean:
	docker ps -aq -f name=$(APP_NAME) | \
		xargs docker stop || true; \
		docker image rm $(IMAGE) || true; \
    go clean

# Swagger commands:
# * Use swagger-setup to install package the first time.
# * Use swagger to refresh the file upon changes to the annotations.
# * Navigate to http://localhost:8080/swagger/index.html

.PHONY: swagger-setup
swagger-setup:
	go get -u github.com/swaggo/swag/cmd/swag && \
		swag init \

.PHONY: swagger
swagger:
	swag init