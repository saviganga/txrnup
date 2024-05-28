package home

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saviganga/txrnup/handlers/home"
)

func WelcomeRoute(app *fiber.App) {
	app.Get("/", home.Home)
}
