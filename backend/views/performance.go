package views

import (
	"strconv"
	"theatre/models"

	"github.com/gofiber/fiber/v2"
)

func PerformanceRoute() *fiber.App {
	performances := fiber.New()
	performances.Mount("hall/", HallRoute())
	performances.Get("/list_performance", func(c *fiber.Ctx) error {
		list_performance := models.ListPerformance{}
		array_list_performance := list_performance.Select()

		list_performance_data := map[int]models.ListPerformanceData{}
		for _, list_p := range array_list_performance {
			list_performance_data[list_p.Id_list_performance] = list_p
		}
		return c.JSON(list_performance_data)
	})

	performances.Get("/type_performance", func(c *fiber.Ctx) error {
		type_performance := models.TypePerformance{}
		array_type_performance := type_performance.Select()

		type_performance_data := map[int]models.TypePerformanceData{}
		for _, type_p := range array_type_performance {
			type_performance_data[type_p.Id_type_performance] = type_p
		}
		return c.JSON(type_performance_data)
	})

	performances.Get("/", func(c *fiber.Ctx) error {

		performance := models.Performance{}
		array_performances := performance.Select()

		performance_data := map[int]models.PerformanceData{}
		for _, perf := range array_performances {
			performance_data[perf.Id_performance] = perf
		}
		return c.JSON(performance_data)
	})

	performances.Get("/:id_performance", func(c *fiber.Ctx) error {

		id_performance := c.Params("id_performance")
		id, err := strconv.Atoi(id_performance)
		if err != nil {
			return c.SendString("Not found performance")
		} else {
			performance := models.Performance{}
			array_performance := performance.Select()
			if len(array_performance)-1 < id {
				return c.SendString("Not found performance")
			} else {
				performance_data := map[int]models.PerformanceData{array_performance[id].Id_performance: array_performance[id]}

				return c.JSON(performance_data)
			}

		}

	})

	return performances
}
