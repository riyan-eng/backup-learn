package util

import "github.com/cloudwego/hertz/pkg/app"

func CurrentUser(c *app.RequestContext) *AccessTokenClaims {
	a, ok := c.Get("claim")
	if !ok {
		panic("can't get user claim")
	}

	return a.(*AccessTokenClaims)
}
