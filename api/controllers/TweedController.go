package controllers

import (
	"FullStackReview/api/models"
	"fmt"
	"net/http"
	"math"
	"strconv"

	"github.com/TwiN/go-away"
	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TweedController struct{
	Router *fiber.App
}

func (t *TweedController) Init(){
	t.Router = fiber.New()

	t.Router.Post("/", t.PostTweed)
	t.Router.Get("/", t.GetTweeds)
	t.Router.Get("/:id", t.GetTweed)
	t.Router.Get("/tweed-range/:skip/:limit/:sort", t.GetTweedsByRange)
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

func (t *TweedController) GetTweedsByRange(c *fiber.Ctx) error{
	tweeds          := []models.Tweed{}
	sort            := c.Params("sort")
	sortValue       := 1
	skip, skipErr   := strconv.Atoi(c.Params("skip", ""))
	limit, limitErr := strconv.Atoi(c.Params("limit", ""))

	if skipErr != nil{
		return c.Status(http.StatusBadRequest).JSON(skipErr)
	}

	if limitErr != nil {
		return c.Status(http.StatusBadRequest).JSON(limitErr)
	}

	if sort == "asc"{
		sortValue = 1
	}else{
		sortValue = -1
	}

	countOpts := options.Count()
 	count, err := mgm.Coll(&models.Tweed{}).CountDocuments(mgm.Ctx(), bson.D{}, countOpts)

	if err != nil{
		return c.Status(http.StatusInternalServerError).JSON(err)
	}

	skip, limit = int(math.Abs(float64(skip))), int(math.Abs(float64(limit)))
	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: sortValue}}).SetLimit(int64(limit)).SetSkip(int64(skip))
	result, err := mgm.Coll(&models.Tweed{}).Find(mgm.Ctx(), bson.D{}, opts)

	if err != nil{
		return c.Status(http.StatusInternalServerError).JSON(err)
	}

 	if err := result.All(mgm.Ctx(), &tweeds); err != nil{
		return c.Status(http.StatusInternalServerError).JSON(err)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"tweeds": tweeds,
		"meta": fiber.Map{
			"total": count,
			"skip": skip,
			"limit": limit,
			"has_more": count > (int64(limit) + int64(skip)),
		},
	})
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