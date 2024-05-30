package xusers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saviganga/txrnup/initialisers"
	"github.com/saviganga/txrnup/models/xuser"
)

func CreateUsers(c *fiber.Ctx) error {
	// fmt.Println("Now i'm creating users")
	db := initialisers.ConnectDb().Db
	user := new(xuser.Xuser)
	err := c.BodyParser(user)
	err = db.Create(&user).Error

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(200).JSON(user)
}

func GetUsers(c *fiber.Ctx) error {
	db := initialisers.ConnectDb().Db
	users := []xuser.Xuser{}
	db.Model(&xuser.Xuser{}).Find(&users).Order("created_at desc")
	return c.Status(200).JSON(users)
}
