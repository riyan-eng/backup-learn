package api

import (
	"fmt"
	"gohtml/infrastructure"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
    username:= c.FormValue("username")
    password:= c.FormValue("password")

    fmt.Println(username, password)
    sess, err:= infrastructure.SessionStore.Get(c)
    if err!=nil{
        panic(err)
    }

    sess.Set("username", username)

    if err := sess.Save(); err != nil {
        panic(err)
    }
    return c.Redirect("/")
}