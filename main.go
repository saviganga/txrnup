package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/saviganga/txrnup/initialisers"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Hello, Ganga! Welcome to Go programming language.")
}

func init() {
	initialisers.LoadEnv()
	initialisers.ConnectDb()
}

func main() {

	port := os.Getenv("PORT")

	app := fiber.New()

	app.Get("/", welcome)
	// fmt.Println(port)

	app.Listen(":" + port)
}
