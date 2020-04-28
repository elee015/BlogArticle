PROJECT_NAME := "blog" 


build-run:
	docker build -t blog-app . && docker-compose up

run:
	docker run -ti -p 8080:8080 --rm --name blog-api blog-app

dev:
	go run main.go

test:
	@go test -v ./...

