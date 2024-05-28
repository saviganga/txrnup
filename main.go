package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/saviganga/txrnup/initialisers"
	"github.com/saviganga/txrnup/routes/home"
)


func init() {
	initialisers.LoadEnv()
	initialisers.ConnectDb()
}

func main() {

	port := os.Getenv("PORT")

	app := fiber.New()

	home.WelcomeRoute(app)


	app.Listen(":" + port)
}
