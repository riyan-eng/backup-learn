package api

import (
	"gohtml/infrastructure"

	"github.com/gofiber/fiber/v2"
)

func Logout(c *fiber.Ctx) error {
    sess, err:= infrastructure.SessionStore.Get(c)
    if err!=nil{
        panic(err)
    }

    if err := sess.Destroy(); err != nil {
        panic(err)
    }
    return c.Redirect("/login")
}