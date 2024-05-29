package xuser

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/saviganga/txrnup/handlers/xuser"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func GetUsers(app *fiber.App) {

	version := os.Getenv("VERSION")
	pathPrefix := fmt.Sprintf("/api/%v/user-service/", version)
	userRoutes := app.Group(pathPrefix, logger.New())

	userRoutes.Get("", xuser.GetUsers)
	userRoutes.Post("", xuser.CreateUsers)


	_ = userRoutes
}
