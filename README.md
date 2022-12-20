# CRUD-golang

I will be using this repository to add everything I will be learning in Golang.

Initially it has 4 endpoints which are:

### POST
##### /v1/user/create
allows to create a user in the database.
JSON: 
{
    "name" : "user_name",
    "age" : 26,
    "address" : {
        "state" : "Col",
        "city" : "Med",
        "pincode" : 5657
    }
}

### GET
#### /v1/user/get
Allows to consult all existing users in the database.

#### /v1/user/get/<<user_name>>
Allows you to query a specific user in the database.

### PUT
#### /v1/user/update
receives a JSON with the same structure as the one in the POST

### DELETE
#### /v1/user/delete/<<<user_name>>
