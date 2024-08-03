package api

import (
	"context"
	"gohertz/infrastructure"
	"gohertz/internal/dto"
	"gohertz/internal/entity"
	"gohertz/util"

	"github.com/cloudwego/hertz/pkg/app"
)

// @Summary     Login
// @Tags       	Authentication
// @Accept		json
// @Produce		json
// @Param       body	body  dto.AuthLogin	true  "body"
// @Router		/auth/login/ [post]
func (m *ServiceServer) AuthLogin(ctx context.Context, c *app.RequestContext) {
	payload := new(dto.AuthLogin)

	if err := c.Bind(payload); err != nil {
		util.NewResponse(c).Error(err.Error(), "", 400)
		return
	}

	errors, errT := util.NewValidation().ValidateStruct(*payload)
	if errT != nil {
		util.NewResponse(c).Error(errors, infrastructure.Localize("FAILED_VALIDATION"), 400)
		return
	}
	user, token, err := m.authService.Login(&ctx, &entity.ServAuthLogin{
		Email:    &payload.Email,
		Password: &payload.Password,
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
		"user": map[string]any{
			"email":     user.Email,
			"username":  user.Username,
			"role_name": user.RoleName,
		},
	}
	util.NewResponse(c).Success(data, nil, infrastructure.Localize("OK_READ"))
}
