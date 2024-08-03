package infrastructure

import "github.com/gofiber/fiber/v2/middleware/session"

var SessionStore *session.Store

func NewSession()  {
    SessionStore = session.New()
}