FRONT_END_BINARY=frontApp
BLOG_BINARY=blogApp
AUTH_BINARY=authApp

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

ps:
	docker-compose ps

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_blog build_auth
	@echo "Stopping docker images (if running...)"
	docker-compose stop blog-service authentication-service
	docker-compose rm -f
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d blog-service authentication-service
	@echo "Docker images built and started!"

up_blog: build_blog
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d blog-service
	@echo "Docker images built and started!"

up_auth: build_auth
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d authentication-service
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

down_blog:
	@echo "Stopping docker images (if running...)"
	docker-compose stop blog-service
	docker-compose rm -f

down_auth:
	@echo "Stopping docker images (if running...)"
	docker-compose stop authentication-service
	docker-compose rm -f

## build_blog: builds the blog binary as a linux executable
build_blog:
	@echo "Building blog binary..."
	cd ../blog-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BLOG_BINARY} ./cmd/api
	@echo "Done!"

## build_auth: builds the blog binary as a linux executable
build_auth:
	@echo "Building auth binary..."
	cd ../authentication-service && env GOOS=linux CGO_ENABLED=0 go build -o ${AUTH_BINARY} ./cmd/api
	@echo "Done!"

## build_front: builds the frone end binary
build_front:
	@echo "Building front end binary..."
	cd ../front-end && env CGO_ENABLED=0 go build -o ${FRONT_END_BINARY} ./cmd/web
	@echo "Done!"

## start: starts the front end
start: build_front
	@echo "Starting front end"
	cd ../front-end && ./${FRONT_END_BINARY} &

## stop: stop the front end
stop:
	@echo "Stopping front end..."
	@-pkill -SIGTERM -f "./${FRONT_END_BINARY}"
	@echo "Stopped front end!"