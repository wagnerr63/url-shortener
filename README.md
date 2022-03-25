# URL Shortener

## Description 
This project was made based on the [URL Shortener challenge](https://github.com/backend-br/desafios/tree/master/1%20-%20Easy/Encurtador%20de%20URL) posted by [Backend-BR](https://github.com/backend-br)

## Requirements
- Golang
- Docker
- Goose (migrations tool)

## How to setup and run the project
- Run the command `docker-compose up -d` to setup the database
- Run `make mig-up` to create the database tables
- Run `go run main.go` to serve the application
