package service

import (
	"context"
	"database/sql"
	"fmt"
	"gohertz/env"
	"gohertz/internal/datastruct"
	"gohertz/internal/entity"
	"gohertz/internal/model"
	"gohertz/internal/repository"
	"gohertz/util"
	"time"

	"github.com/google/uuid"
)

type AuthService interface {
	Register(ctx *context.Context, ent *entity.ServAuthRegister) *util.Error
	Login(ctx *context.Context, ent *entity.ServAuthLogin) (*datastruct.AuthLoginData, *datastruct.AuthToken, *util.Error)
	Refresh(ctx *context.Context, ent *entity.ServAuthRefresh) (*datastruct.AuthToken, *util.Error)
	ResetPassword(ctx *context.Context, ent *entity.ServAuthResetPassword) *util.Error
	ResetPasswordToken(ctx *context.Context, ent *entity.ServAuthResetPasswordToken) *util.Error
	ResetPasswordTokenValidate(ctx *context.Context, ent *entity.ServAuthResetPasswordTokenValidate) *util.Error
	Me(ctx *context.Context, ent *entity.ServAuthMe) (*datastruct.AuthMe, *util.Error)
	Logout(ctx *context.Context, ent *entity.ServAuthLogout) *util.Error
}

type authService struct {
	dao repository.DAO
}

func NewAuthService(dao *repository.DAO) AuthService {
	return &authService{
		dao: *dao,
	}
}

func (m *authService) Register(ctx *context.Context, ent *entity.ServAuthRegister) *util.Error {
	hashPassword, errT := util.GenerateHash(ent.Password)
	if errT != nil {
		return &util.Error{
			Errors: errT.Error(),
		}
	}
	modelUser := model.User{
		Id:       *ent.UserId,
		Email:    sql.NullString{String: *ent.Email, Valid: util.NewIsValid().String(ent.Email)},
		Username: sql.NullString{String: *ent.UserName, Valid: util.NewIsValid().String(ent.UserName)},
		Password: sql.NullString{String: hashPassword, Valid: true},
		IsActive: sql.NullBool{Bool: true, Valid: true},
	}

	modelUserData := model.UserData{
		Id:       uuid.NewString(),
		UserId:   sql.NullString{String: modelUser.Id, Valid: true},
		RoleCode: sql.NullString{String: *ent.RoleCode, Valid: true},
	}

	if err := m.dao.NewAuthRepository().Register(ctx, &modelUser, &modelUserData); err.Errors != nil {
		// custom err
		return err
	}

	return &util.Error{}
}

func (m *authService) Login(ctx *context.Context, ent *entity.ServAuthLogin) (*datastruct.AuthLoginData, *datastruct.AuthToken, *util.Error) {
	token := new(datastruct.AuthToken)

	data, err := m.dao.NewAuthRepository().Login(ctx, ent.Email)
	if err.Errors != nil {
		return data, token, err
	}

	if !data.IsActive {
		return data, token, &util.Error{
			Errors:     "user",
			Message:    "user tidak aktif",
			StatusCode: 400,
		}
	}

	// verify password
	if !util.VerifyHash(data.Password, *ent.Password) {
		return data, token, &util.Error{
			Errors:     "password",
			Message:    "password tidak cocok",
			StatusCode: 400,
		}
	}

	accessToken, accessExpire, err := util.NewToken().CreateAccess(ctx, &data.Id, &data.RoleCode)
	if err.Errors != nil {
		return data, token, &util.Error{
			Errors: err.Errors,
		}
	}
	refreshToken, refreshExpired, err := util.NewToken().CreateRefresh(ctx, &data.Id, &data.RoleCode)
	if err.Errors != nil {
		return data, token, &util.Error{
			Errors: err.Errors,
		}
	}
	return data, &datastruct.AuthToken{
		AccessToken:    accessToken,
		AccessExpired:  accessExpire,
		RefreshToken:   refreshToken,
		RefreshExpired: refreshExpired,
	}, &util.Error{}
}

func (m *authService) Refresh(ctx *context.Context, ent *entity.ServAuthRefresh) (*datastruct.AuthToken, *util.Error) {
	newRefresh := new(datastruct.AuthToken)
	claim, err := util.NewToken().ParseRefresh(ent.Token)
	if err.Errors != nil {
		return newRefresh, &util.Error{
			Errors:     err.Errors,
			StatusCode: 401,
		}
	}

	if err := util.NewToken().ValidateRefresh(ctx, claim); err.Errors != nil {
		return newRefresh, &util.Error{
			Errors:     err.Errors,
			StatusCode: 401,
		}
	}

	accessToken, accessExpire, err := util.NewToken().CreateAccess(ctx, &claim.UserId, &claim.RoleCode)
	if err.Errors != nil {
		return newRefresh, &util.Error{
			Errors: err.Errors,
		}
	}
	refreshToken, refreshExpired, err := util.NewToken().CreateRefresh(ctx, &claim.UserId, &claim.RoleCode)
	if err.Errors != nil {
		return newRefresh, &util.Error{
			Errors: err.Errors,
		}
	}

	return &datastruct.AuthToken{
		AccessToken:    accessToken,
		AccessExpired:  accessExpire,
		RefreshToken:   refreshToken,
		RefreshExpired: refreshExpired,
	}, &util.Error{}
}

func (m *authService) ResetPassword(ctx *context.Context, ent *entity.ServAuthResetPassword) *util.Error {
	claim, err := util.NewToken().ParseReset(ent.Token)
	if err.Errors != nil {
		return &util.Error{
			Errors:     err.Errors,
			StatusCode: 401,
		}
	}

	if err := util.NewToken().ValidateReset(ctx, claim); err.Errors != nil {
		return &util.Error{
			Errors:     err.Errors,
			StatusCode: 401,
		}
	}

	hashPassword, errT := util.GenerateHash(ent.Password)
	if errT != nil {
		return &util.Error{
			Errors: errT.Error(),
		}
	}

	modelUser := model.User{
		Id:       claim.UserId,
		Password: sql.NullString{String: hashPassword, Valid: true},
	}

	if err := m.dao.NewAuthRepository().ResetPassword(ctx, &modelUser); err.Errors != nil {
		return err
	}

	return &util.Error{}
}

func (m *authService) ResetPasswordToken(ctx *context.Context, ent *entity.ServAuthResetPasswordToken) *util.Error {
	data, err := m.dao.NewAuthRepository().Login(ctx, ent.Email)
	if err.Errors != nil {
		return err
	}

	if !data.IsActive {
		return &util.Error{
			Errors:     "user",
			Message:    "user tidak aktif",
			StatusCode: 400,
		}
	}

	resetToken, resetExpire, err := util.NewToken().CreateReset(ctx, &data.Id)
	if err.Errors != nil {
		return &util.Error{
			Errors: err.Errors,
		}
	}

	loc, errT := time.LoadLocation("Asia/Jakarta")
	if errT != nil {
		return &util.Error{Errors: errT.Error()}
	}
	expired := resetExpire.In(loc).Format("15:04:05 02-01-2006")
	link := fmt.Sprintf("%s/atur-ulang-password/?token=%s", env.NewEnv().SERVER_HOST_FE, *resetToken)

	if errT := util.NewSender().ResetPasswordToken(*ent.Email, link, expired); errT != nil {
		return &util.Error{
			Errors: errT.Error(),
		}
	}
	return &util.Error{}
}

func (m *authService) ResetPasswordTokenValidate(ctx *context.Context, ent *entity.ServAuthResetPasswordTokenValidate) *util.Error {
	claim, err := util.NewToken().ParseReset(ent.Token)
	if err.Errors != nil {
		return &util.Error{
			Errors:     err.Errors,
			StatusCode: 401,
		}
	}

	if err := util.NewToken().ValidateReset(ctx, claim); err.Errors != nil {
		return &util.Error{
			Errors:     err.Errors,
			StatusCode: 401,
		}
	}

	return &util.Error{}
}

func (m *authService) Logout(ctx *context.Context, ent *entity.ServAuthLogout) *util.Error {
	err := m.dao.NewAuthRepository().Logout(ctx, ent.UserId)
	if err.Errors != nil {
		return err
	}

	return &util.Error{}
}

func (m *authService) Me(ctx *context.Context, ent *entity.ServAuthMe) (*datastruct.AuthMe, *util.Error) {
	data, err := m.dao.NewAuthRepository().Me(ctx, ent.UserId)
	if err.Errors != nil {
		return data, err
	}

	return data, &util.Error{}
}
