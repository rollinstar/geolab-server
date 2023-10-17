package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rollinstar/geolab-server/database"
	geolab_core "github.com/rollinstar/geolab-server/geolab-core"

	"github.com/rollinstar/geolab-server/model"
)

func GetUser(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")
	fmt.Println("Salted Hash", id)
	var user model.User
	db.Table("member.user").Find(&user, "user_id = ?", id)
	if user.UserId == 0 {
		return c.Status(404).JSON(fiber.Map{})
	}
	return c.Status(200).JSON(fiber.Map{"data": user})
}

func CreateUser(c *fiber.Ctx) error {
	db := database.DB.Db
	type combined struct {
		model.User
		model.Password
	}
	payload := new(combined)
	err := c.BodyParser(payload)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{})
	}

	pwd := payload.Password.Password
	saltedPassword, err := geolab_core.HashAndSalt(pwd)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{})
	}

	err = db.Table("member.user").Create(&payload.User).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{})
	}
	payload.Password.UserId = payload.User.UserId
	payload.Password.Password = saltedPassword

	err = db.Table("auth.password").Create(&payload.Password).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{})
	}
	return c.Status(201).JSON(fiber.Map{"data": payload})
}
