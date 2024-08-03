package api

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func (m *ServiceServer) Index(c context.Context, ctx *app.RequestContext) {
	ctx.JSON(200, utils.H{
		"message": "Welcome to my channel",
	})
}
