package mysqldatabase

import (
	"fmt"
	"gin_gonic_api/pkg/controllers"
	"gin_gonic_api/pkg/platform/configuration"
	"gin_gonic_api/pkg/services"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() controllers.PostController {
	configuration.LoadEnvVariables()
	DB := ConnectToSQL()

	fmt.Println("MYSQL connection estabilished")

	postservice := services.NewPostService(DB)
	postcontroller := controllers.NewPost(postservice)

	return postcontroller
}

func ConnectToSQL() *gorm.DB {
	connectionString := os.Getenv("DB_URL")
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	return db
}
