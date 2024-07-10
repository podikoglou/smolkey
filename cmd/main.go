package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/podikoglou/smolkey/internal/engine"
)

func main() {
	// create engine instance
	engine := *engine.NewEngine()

	// create fiber app for serving	the store
	app := fiber.New()

	app.Put("/:key", func(c fiber.Ctx) error {
		key := c.Params("key")
		value := c.FormValue("value", "")

		// put key in the store
		err := engine.Put(key, []byte(value))

		// handle errors
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		// return http code 201 "created"
		return c.SendStatus(201)
	})

	app.Get("/:key", func(c fiber.Ctx) error {
		key := c.Params("key")

		// try to look up key in the store
		value, err := engine.Get(key)

		// handle rrors
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.Send(value)
	})

	log.Fatal(app.Listen(":8080"))
}
