package home

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/saviganga/txrnup/handlers/home"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func WelcomeRoute(app *fiber.App) {

	version := os.Getenv("VERSION")
	pathPrefix := fmt.Sprintf("/api/%v/welcome-service/", version)
	welcomeRoutes := app.Group(pathPrefix, logger.New())

	welcomeRoutes.Get("", home.Home)

	_ = welcomeRoutes
}
