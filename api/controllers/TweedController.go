package controllers

import (
	"FullStackReview/api/models"
	"fmt"
	"net/http"


	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/TwiN/go-away"
)

type TweedController struct{
	Router *fiber.App
}

func (t *TweedController) Init(){
	t.Router = fiber.New()

	t.Router.Get("/", t.GetTweeds)
	t.Router.Get("/:id", t.GetTweed)
	t.Router.Post("/", t.PostTweed)
}

func (t *TweedController) GetTweed(c *fiber.Ctx) error {
	id := c.Params("id")
	tweed := models.Tweed{}

	fmt.Println("id:", id)
	if err := mgm.Coll(&models.Tweed{}).FindByID(id, &tweed); err != nil{
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	return c.Status(http.StatusOK).JSON(tweed)
}

func (t *TweedController) GetTweeds(c *fiber.Ctx) error{
	tweeds := []models.Tweed{}

	result, err := mgm.Coll(&models.Tweed{}).Find(mgm.Ctx(), bson.D{})

	if err != nil{
		return c.Status(http.StatusInternalServerError).JSON(err)
	}

 	if err := result.All(mgm.Ctx(), &tweeds); err != nil{
		return c.Status(http.StatusInternalServerError).JSON(err)
	}

	return c.Status(http.StatusOK).JSON(tweeds[len(tweeds) - 10:])
}

func (t *TweedController) PostTweed(c *fiber.Ctx) error{
	tweed := models.Tweed{}

	if err := c.BodyParser(&tweed); err != nil{
		return c.Status(http.StatusBadRequest).JSON(err)
	}
	
	if err := tweed.Validate(); err != nil{
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	tweed.Name    = goaway.Censor(tweed.Name)
	tweed.Content = goaway.Censor(tweed.Content)

	fmt.Println(tweed)
	if err := mgm.Coll(&models.Tweed{}).Create(&tweed); err != nil{
		return c.Status(http.StatusInternalServerError).JSON(err)
	}

	return c.Status(http.StatusOK).JSON(tweed)
}