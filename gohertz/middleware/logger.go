package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/requestid"
)

// LoggerMiddleware middleware for logging incoming requests
func Logger() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		start := time.Now()
		defer func() {
			stop := time.Now()

			rid := requestid.Get(c)
			fmt.Println("receiver:", rid)

			log_message := fmt.Sprintf("%v | %v | %v | %v | %v", start.Format("15:04 MST"), c.Response.StatusCode(), stop.Sub(start).String(), string(c.Method()), string(c.Path()))
			hlog.Info(log_message)

		}()
		c.Next(ctx)
	}
}
