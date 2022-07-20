package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/imayrus/url-shortener/handlers"
)

func SetupAndListen() {

	r := fiber.New()

	r.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	r.Get("/:url", handlers.Redirect)
	r.Post("shorturl", handlers.CreateUrl)

	r.Listen(":3000")

}
