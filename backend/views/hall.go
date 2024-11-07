package views

import (
	"fmt"
	"strconv"
	"theatre/models"

	"github.com/gofiber/fiber/v2"
)

func HallRoute() *fiber.App {
	halls := fiber.New()

	halls.Get("/", func(c *fiber.Ctx) error {

		hall := models.Hall{}
		array_halls := hall.Select()

		hall_data := map[int]models.HallData{}
		for _, h := range array_halls {
			hall_data[h.Id_hall] = h
		}
		return c.JSON(hall_data)
	})

	halls.Get("/:id_theatre", func(c *fiber.Ctx) error {

		id_theatre := c.Params("id_theatre")
		id, err := strconv.Atoi(id_theatre)
		if err != nil {
			return c.SendString("Not found theatre")
		} else {
			hall := models.Hall{}
			query := fmt.Sprintf("SELECT * FROM Hall WHERE id_theatre = %d", id)
			array_halls := hall.SelectQuery(query)
			hall_data := map[int]models.HallData{}

			if len(array_halls) <= 0 {
				return c.SendString("Not found hall")
			} else {

				for _, h := range array_halls {
					hall_data[h.Id_hall] = h
				}
				return c.JSON(hall_data)
			}
		}
	})

	halls.Get("/rows/:id_hall", func(c *fiber.Ctx) error {

		id_hall := c.Params("id_hall")
		id, err := strconv.Atoi(id_hall)
		if err != nil {
			return c.SendString("Not found hall")
		} else {
			row := models.Row{}
			array_rows := row.Select()
			row_data := map[int]models.RowData{}

			if len(array_rows) <= 0 {
				return c.SendString("Not found rows")
			} else {
				for _, r := range array_rows {
					if(id == r.Id_hall){
						row_data[r.Id_row] = r
					}
				}
				return c.JSON(row_data)
			}
		}
	})

	halls.Get("/rows/places/:id_row", func(c *fiber.Ctx) error {
		id_row := c.Params("id_row")
		id, err := strconv.Atoi(id_row)
		if err != nil {
			return c.SendString("Not found row")
		} else {
			place := models.Place{}
			array_places := place.Select()
			place_data := map[int]models.PlaceData{}

			if len(array_places)-1 < id {
				return c.SendString("Not found places")
			} else {
				for _, r := range array_places {
					place_data[r.Id_place] = r
				}
				return c.JSON(place_data)
			}
		}
	})

	halls.Get("/rows/places/ticket/:id_place", func(c *fiber.Ctx) error {
		id_place := c.Params("id_place")
		id, err := strconv.Atoi(id_place)
		if err != nil {
			return c.SendString("Not found place")
		} else {
			ticket := models.Ticket{}
			query := fmt.Sprintf(`SELECT * FROM Ticket WHERE id_place = %d`, id)
			array_tickets := ticket.SelectQuery(query)
			if len(array_tickets) > 0 {
				return c.SendString("1")
			} else {
				return c.SendString("0")
			}
		}
	})

	return halls
}
