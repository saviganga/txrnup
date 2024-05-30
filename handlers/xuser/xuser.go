package xuser

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saviganga/txrnup/controllers/xusers"
)

func GetUsers(c *fiber.Ctx) error {
	return xusers.GetUsers(c)

}

func CreateUsers(c *fiber.Ctx) error {
	return xusers.CreateUsers(c)
}
