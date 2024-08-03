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
// @Param       body	body  dto.AuthResetPasswordToken	true  "body"
// @Router		/auth/reset-password-token/ [post]
func (m *ServiceServer) AuthResetPasswordToken(ctx context.Context, c *app.RequestContext) {
	payload := new(dto.AuthResetPasswordToken)

	if err := c.Bind(payload); err != nil {
		util.NewResponse(c).Error(err.Error(), "", 400)
		return
	}

	errors, errT := util.NewValidation().ValidateStruct(*payload)
	if errT != nil {
		util.NewResponse(c).Error(errors, infrastructure.Localize("FAILED_VALIDATION"), 400)
		return
	}
	if err := m.authService.ResetPasswordToken(&ctx, &entity.ServAuthResetPasswordToken{
		Email: &payload.Email,
	}); err.Errors != nil {
		util.NewResponse(c).Error(err.Errors, err.Message, err.StatusCode)
		return
	}

	util.NewResponse(c).Success(nil, nil, infrastructure.Localize("OK_READ"))
}
