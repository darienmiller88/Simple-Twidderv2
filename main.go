package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
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
	app.Use(limiter.New(limiter.Config{
		Max:            2000,
		Expiration:     24 * time.Hour,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"message": "You reached your limit of 20 tweeds per day. Come back tomorrow!",
			})
		},
	}))

	tweedController.Init()
	app.Mount("/api/v1/tweeds", tweedController.Router)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "hello world"})
	})
	
	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}