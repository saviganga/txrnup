package xuser

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/saviganga/txrnup/controllers/xusers"
)

func GetUsers(c *fiber.Ctx) error {
	return c.SendString("Now i'm gonna try to fetch users")
}

func CreateUsers(c *fiber.Ctx) error {
	fmt.Println("Now i'm gonna try to create users")

	return xusers.CreateUsers(c)
}