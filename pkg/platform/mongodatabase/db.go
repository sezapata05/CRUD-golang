package mongodatabase

import (
	"context"
	"fmt"

	"gin_gonic_api/pkg/controllers"
	"gin_gonic_api/pkg/platform/mongodatabase/configuration"
	"gin_gonic_api/pkg/services"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	userservice    services.UserService
	usercontroller controllers.UserController
	ctx            context.Context
	usercollection *mongo.Collection
	// mongoclient    *mongo.Client
)

func Init() (controllers.UserController, context.Context) {
	config, err := configuration.LoadConfig("configuration.yml")
	if err != nil {
		log.Fatal(err)
	}

	ctx = context.TODO()

	// mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoconn := options.Client().ApplyURI(fmt.Sprintf("%v://%v:%v", config.DB.Typedb, config.DB.Host, config.DB.Port))
	mongoclient, err := mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal(err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mongo connection estabilished")

	usercollection = mongoclient.Database("userdb").Collection("user")
	userservice = services.NewUserService(usercollection, ctx)
	usercontroller = controllers.New(userservice)

	return usercontroller, ctx
}
