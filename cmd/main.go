package main

import (
	"gin_gonic_api/pkg/platform/mongodatabase"
	"gin_gonic_api/pkg/platform/mysqldatabase"
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
	postcontroller := mysqldatabase.Init()

	defer mongoclient.Disconnect(ctx)

	server = gin.Default()

	basepath := server.Group("/v1")

	usercontroller.RegisterUserRoutes(basepath)

	postcontroller.RegisterPostRoutes(basepath)

	log.Fatal(server.Run())

}
