package handler

import (
	"github.com/Winens/Wisp-Server/db"
	"github.com/Winens/Wisp-Server/model"
	"github.com/Winens/Wisp-Server/pkg/session"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

var validate = validator.New()

func SignUp(c *fiber.Ctx) error {
	sess, err := session.Store.Get(c)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var form struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&form); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if err := validate.Struct(form); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	var user model.User
	db.PostgeSQL.Where("username = ?", form.Username).Limit(1).First(&user)
	if user.ID != 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "Username already exists",
		})
	}

	if err := user.GeneratePasswordDigest(form.Password); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	user.Username = form.Username

	authId, err := gonanoid.New(48)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	user.AuthID = authId

	db.PostgeSQL.Create(&user)

	sess.Set("authId", authId)
	return c.SendStatus(fiber.StatusOK)
}
