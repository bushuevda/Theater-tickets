package views

import (
	"strconv"
	"theatre/models"
	"github.com/gofiber/fiber/v2"
)

func TheatreRoute() *fiber.App {
	theatres := fiber.New()
	theatres.Mount("performance/", PerformanceRoute())
	
	theatres.Get("/locality", func(c *fiber.Ctx) error {
		locality := models.Locality{}
		array_localities := locality.Select()

		locality_data := map[int]models.LocalityData{}
		for _, loc := range array_localities {
			locality_data[loc.Id_locality] = loc
		}
		return c.JSON(locality_data)
	})

	theatres.Get("/", func(c *fiber.Ctx) error {

		theatre := models.Theatre{}
		array_theatres := theatre.Select()

		theatre_data := map[int]models.TheatreData{}
		for _, th := range array_theatres {
			theatre_data[th.Id_theatre] = th
		}
		return c.JSON(theatre_data)
	})

	theatres.Get("/:id_theatre", func(c *fiber.Ctx) error {

		id_theatre := c.Params("id_theatre")
		id, err := strconv.Atoi(id_theatre)
		if err != nil {
			return c.SendString("Not found theatre")
		} else {
			theatre := models.Theatre{}
			array_theatres := theatre.Select()
			if len(array_theatres)-1 < id {
				return c.SendString("Not found theatre")
			} else {
				theatre_data := map[int]models.TheatreData{array_theatres[id].Id_theatre: array_theatres[id]}

				return c.JSON(theatre_data)
			}

		}

	})

	return theatres
}
