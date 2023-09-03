package middleware

import (
	"github.com/Winens/Wisp-Server/db"
	"github.com/Winens/Wisp-Server/model"
	"github.com/Winens/Wisp-Server/pkg/session"
	"github.com/gofiber/fiber/v2"
)

func MustAuthenticate(c *fiber.Ctx) error {
	sess, err := session.Store.Get(c)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var user model.User
	db.PostgeSQL.Where("auth_id = ?", sess.Get("authId")).Limit(1).First(&user)

	if user.ID == 0 {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	c.Locals("sessionUser", user)
	return c.Next()
}
