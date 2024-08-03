package dto

type AuthRegister struct {
	Email    string `json:"email" valid:"required;email"`
	UserName string `json:"username"`
	Password string `json:"password" valid:"required;min:8"`
	RoleCode string `json:"role_code" valid:"required;in:STAFF,USER"`
}

type AuthLogin struct {
	Email    string `json:"email" valid:"required;email"`
	Password string `json:"password" valid:"required"`
}

type AuthRefresh struct {
	Token string `json:"token" valid:"required"`
}

type AuthResetPasswordToken struct {
	Email string `json:"email" valid:"required;email"`
}
