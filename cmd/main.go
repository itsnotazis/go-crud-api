package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/itsnotazis/go-crud-api/cmd/api"
)

func main() {
	f := fiber.New()

	app, err := api.NewAPIServer(f, &api.APIConfig{
		AppName:       "Go CRUD API",
		ServerHeader:  "Fiber",
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
	}, ":8080")

	if err != nil {
		log.Fatal(err)
	}

	if err := app.Listen(); err != nil {
		log.Fatal(err)
	}
}
