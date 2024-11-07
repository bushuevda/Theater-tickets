package main

import (
	"fmt"
	models "theatre/models"
	"theatre/views"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/mattn/go-sqlite3"
)

func testDb() {

	mu := models.Performance{}
	for _, b := range mu.Select() {

		fmt.Println(b)
	}

}

func main() {
	testDb()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Mount("theatre/", views.TheatreRoute())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(" there!!")
	})

	app.Get("/df", func(c *fiber.Ctx) error {
		movies := map[int]string{1: "Toy story", 2: "The Raid", 3: "Hero", 4: "Ip Man", 5: "Kung Fu Panda"}

		return c.JSON(movies)
	})

	app.Listen("localhost:8000")

}
