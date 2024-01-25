## Golang HTTP Server

# A simple http server created in golang using net/http package.

# Steps to create HTTP Server:
1) Add Handler functions to handle routes
2) Read data from request
3) Write data to output stream
4) Listen and Server the handlers 

# Steps to run any general golang code
1) Create golang file
   Ex: -> main.go
   
2) Run `go mod init <package_name>`
   Ex: `go mod init example.com/simple-http-server`
   This will create a go.mod file
   
3) Install dependencies using `go mod tidy`
   Ex: `go mod tidy`
   This will create a go.sum file
   
5) Run go file directly using `go run <file_name>`
   Ex: `go run main.go`
   
7) We can also compile code to build executable binary
   Run `go build <file_name>` or `go build .`
   Ex: `go build main.go`
   This will create an executable binary with same name by default

   We can also specify output binary filename
   Run `go build -o <filename> .`
   Ex: `go build -o server .`

   We can directly execute generated binary
   Ex: `./server`

# Dockerize application
Look at Dockerfile for steps to create docker image
