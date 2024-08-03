package api

import (
	"context"
	"gohertz/infrastructure"
	"gohertz/internal/dto"
	"gohertz/internal/entity"
	"gohertz/util"

	"github.com/cloudwego/hertz/pkg/app"
)

// @Summary     Refresh
// @Tags       	Authentication
// @Accept		json
// @Produce		json
// @Param       body	body  dto.AuthRefresh	true  "body"
// @Router		/auth/refresh/ [post]
func (m *ServiceServer) AuthRefresh(ctx context.Context, c *app.RequestContext) {
	payload := new(dto.AuthRefresh)

	if err := c.Bind(payload); err != nil {
		util.NewResponse(c).Error(err.Error(), "", 400)
		return
	}

	errors, errT := util.NewValidation().ValidateStruct(*payload)
	if errT != nil {
		util.NewResponse(c).Error(errors, infrastructure.Localize("FAILED_VALIDATION"), 400)
		return
	}
	token, err := m.authService.Refresh(&ctx, &entity.ServAuthRefresh{
		Token: &payload.Token,
	})
	if err.Errors != nil {
		util.NewResponse(c).Error(err.Errors, err.Message, err.StatusCode)
		return
	}

	data := map[string]any{
		"access_token":    token.AccessToken,
		"access_expired":  token.AccessExpired.Time.Local(),
		"refresh_token":   token.RefreshToken,
		"refresh_expired": token.RefreshExpired.Time.Local(),
	}
	util.NewResponse(c).Success(data, nil, infrastructure.Localize("OK_READ"))
}
