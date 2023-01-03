# CRUD-golang

I will be using this repository to add everything I will be learning in Golang.

Initially it has 4 endpoints which are:

The repository contains 2 types of functionalities (connections) to databases such as MongoDB and MySQL.

Within the endpoints are exposed the groups "user" and "post".

which contain the operations:

## 1 POST
## 2 GET
#### * Get all
#### * Get single
## 1. PUT
## 1 DELETE


you can visit and download the official collections that have been used to work in 
/pgk/doc/Api - go.postman_collection.json

# to run locally you need to create the following files with the following structure:
## configuration.yml
```
Database:
  type_db: mongodb
  host: <<your_host>>
  port: <<your_port>>
```
## .env
```
PORT=<<gin port app>>
DB_URL="user:pass@tcp(host_name:3306)/database_name?parseTime=True"
```

# And run the following command
```
go run cmd/main.go
```
