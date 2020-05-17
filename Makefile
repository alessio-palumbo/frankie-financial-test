APP_NAME   = frankie-financial-test
DOCKERFILE = images/Dockerfile
COMMIT_SHA = $(shell git rev-parse --short HEAD)
IMAGE      = $(APP_NAME):$(COMMIT_SHA)
PORT       = 8080

.PHONY: build
build: clean
	@echo 'Building executable...\n' && \
	go build -o $(APP_NAME) main.go

.PHONY: run-build
run-build: build
	@echo 'Running build...\n' && \
	./$(APP_NAME)

.PHONY: run
run:
	@echo 'Running main.go...\n' && \
	go run main.go

.PHONY: test
test:
	@echo 'Running tests...\n' && \
	go test -v ./...

.PHONY: test-cover
test-cover:
	@echo 'Test coverage:\n' && \
	go test -coverprofile cover.out ./... && \
	go tool cover -func cover.out && \
	rm cover.out

.PHONY: clean
clean:
	@echo 'Cleaning up build...\n' && \
	go clean

# Docker

.PHONY: docker-build
docker-build: docker-clean
	@echo 'Building container...\n' && \
	docker build -t $(IMAGE) -f $(DOCKERFILE) .

.PHONY: docker-run
docker-run: docker-build
	@echo '\nRunning container $(APP_NAME) on port $(PORT)...\n' && \
	docker container run -d --rm -p $(PORT):8080 --name $(APP_NAME) $(IMAGE)

.PHONY: docker-clean
docker-clean:
	@echo 'Cleaning up...\n' && \
	docker ps -aq -f name=$(APP_NAME) | \
	xargs docker stop 2> /dev/null; \
	docker image rm $(IMAGE) 2> /dev/null; \
	go clean

# Swagger commands:
# * Use swagger-setup to install package the first time.
# * Use swagger to refresh the file upon changes to the annotations.
# * Navigate to http://localhost:$(PORT)/swagger/index.html

.PHONY: swagger-setup
swagger-setup:
	@echo '\nDownloading swaggo and generating docs...\n' && \
	go get -u github.com/swaggo/swag/cmd/swag && \
	swag init \

.PHONY: swagger
swagger:
	@echo '\nGenerating swagger docs...\n' && \
	swag init && \
	echo '\ncurl localhost:$(PORT)/swagger/doc.json to see docs.\n'