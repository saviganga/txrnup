package xusers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/saviganga/txrnup/initialisers"
	"github.com/saviganga/txrnup/models/xuser"
)

func CreateUsers(c *fiber.Ctx) error {
	// fmt.Println("Now i'm creating users")
	db := initialisers.ConnectDb().Db
	user := new(xuser.Xuser)
	fmt.Println(len(user.Email))
	err := c.BodyParser(user)
	err = db.Create(&user).Error

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(200).JSON(user)
}
