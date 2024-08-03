package middleware

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/google/uuid"
	"github.com/hertz-contrib/requestid"
)

func RequestId() app.HandlerFunc {
	return requestid.New(requestid.WithGenerator(func(ctx context.Context, c *app.RequestContext) string {
		newRequestID := uuid.NewString()
		fmt.Println("sender:", newRequestID)
		return newRequestID
	}))
}
