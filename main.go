package main

import (
	"fmt"
	"log"

	"github.com/abhishek2966/example-fiber/config"
	"github.com/abhishek2966/example-fiber/handler"
	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	data := config.Data{}
	data.InitFlag()
	yamldata, err := data.ReadYAML()
	if err != nil {
		panic(err)
	}
	err = data.DecodeYAML(yamldata)
	if err != nil {
		log.Print(err)
	}
	addr := fmt.Sprintf(":%v", data.Port)

	r := fiber.New()
	r.Use(handler.TimeLapsedMiddleware)
	r.Get("/photos", handler.HandlePhotosFetch)
	r.Post("/posts", handler.HandlePostsSave)

	r.Listen(addr)
}
