package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/saviganga/txrnup/initialisers"
	"github.com/saviganga/txrnup/routes/home"
	"github.com/saviganga/txrnup/routes/xuser"
)

func init() {
	initialisers.LoadEnv()
	initialisers.ConnectDb()
}

func main() {

	port := os.Getenv("PORT")
	fmt.Println(os.Getenv("VERSION"))

	app := fiber.New()
	app.Use(logger.New())

	home.WelcomeRoute(app)
	xuser.GetUsers(app)

	app.Listen(":" + port)
}
