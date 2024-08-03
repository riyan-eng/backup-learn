package middleware

import (
	"context"
	"gohertz/util"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
)

func Jwt() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		authHeader := c.GetHeader("Authorization")
		bearerString := string(authHeader)
		if bearerString == "" {
			util.NewResponse(c).Error("authorization header is required", "", 400)
			return
		}
		token, found := strings.CutPrefix(bearerString, "Bearer ")
		if !found {
			util.NewResponse(c).Error("undefined token", "", 400)
			return

		}
		claims, err := util.NewToken().ParseAccess(&token)
		if err.Errors != nil {
			util.NewResponse(c).Error(err.Errors, "", 401)
			return
		}
		if err := util.NewToken().ValidateAccess(&ctx, claims); err.Errors != nil {
			util.NewResponse(c).Error(err.Errors, "", 401)
			return
		}

		c.Set("claim", claims)
		c.Next(ctx)
	}
}
