package session

import "github.com/gofiber/fiber/v2/middleware/session"

var Store *session.Store

func Init() {
	Store = session.New(session.Config{})
}
