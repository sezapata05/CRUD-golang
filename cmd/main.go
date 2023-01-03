package main

import (
	// "gin_gonic_api/pkg/platform/mongo/"
	"gin_gonic_api/pkg/platform/mongodatabase"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	server      *gin.Engine
	mongoclient *mongo.Client
)

func main() {

	usercontroller, ctx := mongodatabase.Init()

	// database
	// database.Init()

	defer mongoclient.Disconnect(ctx)

	server = gin.Default()

	basepath := server.Group("/v1")

	usercontroller.RegisterUserRoutes(basepath)

	log.Fatal(server.Run(":9090"))

}
