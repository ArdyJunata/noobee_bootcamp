package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/google/uuid"
)

type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

func main() {
	router := fiber.New()
	users := []User{}

	router.Use(requestid.New())
	router.Use(Trace())

	router.Post("/users", func(c *fiber.Ctx) error {
		req := User{}

		if err := c.BodyParser(&req); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"success":     false,
				"status_code": http.StatusBadRequest,
				"message":     err.Error(),
			})
		}

		req.ID = uuid.NewString()

		users = append(users, req)

		return c.Status(http.StatusCreated).JSON(fiber.Map{
			"success":     true,
			"status_code": http.StatusCreated,
			"message":     "created success",
		})
	})

	router.Get("/users", func(c *fiber.Ctx) error {

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"success":     true,
			"status_code": http.StatusOK,
			"message":     "get all success success",
			"payload":     users,
		})
	})

	router.Put("/users/:id", func(c *fiber.Ctx) error {
		req := User{}
		id := c.Params("id", "")

		if err := c.BodyParser(&req); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"success":     false,
				"status_code": http.StatusBadRequest,
				"message":     err.Error(),
			})
		}

		for i := range users {
			if users[i].ID == id {
				users[i].Name = req.Name
				users[i].Email = req.Email
				users[i].Address = req.Address
				break
			}
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"success":     true,
			"status_code": http.StatusOK,
			"message":     "update success",
		})
	})

	router.Delete("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id", "")

		for i := range users {
			if users[i].ID == id {
				users = append(users[:i], users[i+1:]...)
				break
			}
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"success":     true,
			"status_code": http.StatusOK,
			"message":     "delete success",
		})
	})

	router.Listen(":4444")
}

func Trace() fiber.Handler {
	return func(c *fiber.Ctx) error {
		traceId := string(c.Response().Header.Peek("X-Request-Id"))
		log.Printf("message = incoming request, method = %s, uri = %s, trace_id = %s", c.Method(), c.Path(), traceId)
		err := c.Next()
		log.Printf("message = finish` request, method = %s, uri = %s, trace_id = %s", c.Method(), c.Path(), traceId)
		return err
	}
}
