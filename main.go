package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"

	"FullStackReview/api/controllers"
)

func main(){
	godotenv.Load()
	app := fiber.New()
	err := mgm.SetDefaultConfig(nil, "UserTweeds", options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	tweedController := controllers.TweedController{}

	if err != nil{
		log.Fatal(err)
	}

	app.Use(cors.New())
	app.Use(logger.New())

	tweedController.Init()
	app.Mount("/api/v1/tweeds", tweedController.Router)
	
	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}