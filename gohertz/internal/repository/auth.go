package repository

import (
	"context"
	"database/sql"
	"fmt"
	"gohertz/infrastructure"
	"gohertz/internal/datastruct"
	"gohertz/internal/model"
	"gohertz/util"

	"github.com/blockloop/scan/v2"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type AuthRepository interface {
	Register(ctx *context.Context, modelUser *model.User, modelUserData *model.UserData) *util.Error
	Login(ctx *context.Context, id *string) (*datastruct.AuthLoginData, *util.Error)
	ResetPassword(ctx *context.Context, modelUser *model.User) *util.Error
	Logout(ctx *context.Context, userId *string) *util.Error
	Me(ctx *context.Context, userId *string) (*datastruct.AuthMe, *util.Error)
}

type authRepository struct {
	sqlDB  *sql.DB
	sqlxDB *sqlx.DB
	cache  *redis.Client
}

func (m *authRepository) Register(ctx *context.Context, modelUser *model.User, modelUserData *model.UserData) *util.Error {
	tx := m.sqlxDB.MustBegin()
	_, errT := tx.NamedExecContext(*ctx, `
	insert into users (uuid, email, username, password, is_active) values (:id, :email, :username, :password, :is_active)
	`, modelUser)
	if errT != nil {
		return &util.Error{
			Errors:  errT.Error(),
			Message: infrastructure.Localize("FAILED_CREATE_NO_DATA"),
		}
	}
	_, errT = tx.NamedExecContext(*ctx, `
	insert into user_datas (uuid, user_uuid, role_code) values (:id, :user_id, :role_code)
	`, modelUserData)

	if errT != nil {
		return &util.Error{
			Errors:  errT.Error(),
			Message: infrastructure.Localize("FAILED_CREATE_NO_DATA"),
		}
	}
	if errT := tx.Commit(); errT != nil {
		return &util.Error{
			Errors:  errT.Error(),
			Message: infrastructure.Localize("FAILED_CREATE_NO_DATA"),
		}
	}
	return &util.Error{}
}

func (m *authRepository) Login(ctx *context.Context, email *string) (*datastruct.AuthLoginData, *util.Error) {
	data := new(datastruct.AuthLoginData)

	query := fmt.Sprintf(`
	select u."uuid", u.username, u.email, u."password", u.is_active, r.code as role_code, r."name" as role_name 
	from users u 
	left join user_datas ud on ud.user_uuid = u."uuid" 
	left join roles r on r.code = ud.role_code 
	where u.email = '%v' 
	limit 1
	`, *email)

	sqlRows, err := m.sqlDB.QueryContext(*ctx, query)
	if err != nil {
		return data, &util.Error{
			Errors: err.Error(),
		}
	}

	if err := scan.Row(data, sqlRows); err != nil {
		return data, &util.Error{
			Errors:     err.Error(),
			StatusCode: 400,
			Message:    "tidak terdaftar",
		}
	}

	return data, &util.Error{}
}

func (m *authRepository) ResetPassword(ctx *context.Context, modelUser *model.User) *util.Error {
	sqlRslt, errT := m.sqlxDB.NamedExecContext(*ctx, "update users set password=coalesce(:password, password), updated_at=now() where uuid=:uuid", modelUser)
	if errT != nil {
		return &util.Error{
			Errors:  errT.Error(),
			Message: infrastructure.Localize("FAILED_UPDATE"),
		}
	}
	rowsAffected, errT := sqlRslt.RowsAffected()
	if errT != nil {
		return &util.Error{
			Errors:  errT.Error(),
			Message: infrastructure.Localize("FAILED_UPDATE"),
		}
	}
	if rowsAffected == 0 {
		return &util.Error{
			Errors:  "no row affected",
			Message: infrastructure.Localize("FAILED_UPDATE_NO_DATA"),
		}
	}

	return &util.Error{}
}

func (m *authRepository) Logout(ctx *context.Context, userId *string) *util.Error {
	if err := m.cache.Del(*ctx, fmt.Sprintf("access-token-%s", *userId)).Err(); err != nil {
		return &util.Error{
			Errors: err.Error(),
		}
	}
	return &util.Error{}
}

func (m *authRepository) Me(ctx *context.Context, userId *string) (*datastruct.AuthMe, *util.Error) {
	data := new(datastruct.AuthMe)

	query := fmt.Sprintf(`
	select u."uuid", u.username, u.email, r.code as role_code, r."name" as role_name 
	from users u 
	left join user_datas ud on ud.user_uuid = u."uuid" 
	left join roles r on r.code = ud.role_code 
	where u.uuid = '%v' 
	limit 1
	`, *userId)

	sqlRows, err := m.sqlDB.QueryContext(*ctx, query)
	if err != nil {
		return data, &util.Error{
			Errors: err.Error(),
		}
	}

	if err := scan.Row(data, sqlRows); err != nil {
		return data, &util.Error{
			Errors:     err.Error(),
			StatusCode: 400,
			Message:    "tidak terdaftar",
		}
	}

	return data, &util.Error{}
}
