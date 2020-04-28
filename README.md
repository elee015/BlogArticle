# BlogArticle


## Description: 
RESTful API for articles CR

## To run:
`make build-run` or `docker build -t blog-app . && docker-compose up`

## To test
`make test` or `go test -v ./...`

## API:
- GET "/articles"
- GET "/articles/{Id}"
- POST "/articles"