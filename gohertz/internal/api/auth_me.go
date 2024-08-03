package api

import (
	"context"
	"gohertz/infrastructure"
	"gohertz/internal/entity"
	"gohertz/util"

	"github.com/cloudwego/hertz/pkg/app"
)

// @Summary     Me
// @Tags       	Authentication
// @Accept		json
// @Produce		json
// @Router		/auth/me/ [get]
func (m *ServiceServer) AuthMe(ctx context.Context, c *app.RequestContext) {
	user := util.CurrentUser(c)

	data, err := m.authService.Me(&ctx, &entity.ServAuthMe{UserId: &user.UserId})
	if err.Errors != nil {
		util.NewResponse(c).Error(err.Errors, err.Message, err.StatusCode)
		return
	}

	util.NewResponse(c).Success(data, nil, infrastructure.Localize("OK_READ"))
}
