package api

import (
	"context"
	"gohertz/infrastructure"
	"gohertz/internal/dto"
	"gohertz/internal/entity"
	"gohertz/util"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/google/uuid"
)

// @Summary     Register
// @Tags       	Authentication
// @Accept		json
// @Produce		json
// @Param       body	body  dto.AuthRegister	true  "body"
// @Router		/auth/register/ [post]
func (m *ServiceServer) AuthRegister(ctx context.Context, c *app.RequestContext) {
	payload := new(dto.AuthRegister)

	if err := c.Bind(payload); err != nil {
		util.NewResponse(c).Error(err.Error(), "", 400)
		return
	}

	errors, errT := util.NewValidation().ValidateStruct(*payload)
	if errT != nil {
		util.NewResponse(c).Error(errors, infrastructure.Localize("FAILED_VALIDATION"), 400)
		return
	}
	userId := uuid.NewString()
	if err := m.authService.Register(&ctx, &entity.ServAuthRegister{
		UserId:   &userId,
		Email:    &payload.Email,
		UserName: &payload.UserName,
		Password: &payload.Password,
		RoleCode: &payload.RoleCode,
	}); err.Errors != nil {
		util.NewResponse(c).Error(err.Errors, err.Message, err.StatusCode)
		return
	}

	data := map[string]any{
		"id": userId,
	}
	util.NewResponse(c).Success(data, nil, infrastructure.Localize("OK_CREATE"), 201)
}
