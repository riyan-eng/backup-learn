package util

import (
	"context"
	"encoding/json"
	"fmt"
	"gohertz/env"
	"gohertz/infrastructure"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

var mySigningKey = []byte("AllYourBase")

type AccessTokenClaims struct {
	UserId   string `json:"user_id"`
	RoleCode string `json:"role_code"`
	UUID     string `json:"id"`
	jwt.RegisteredClaims
}

type AccessTokenCached struct {
	AccessUID string `json:"access"`
}

type RefreshTokenClaims struct {
	UserId   string `json:"user_id"`
	RoleCode string `json:"role_code"`
	UUID     string `json:"id"`
	jwt.RegisteredClaims
}

type RefreshTokenCached struct {
	RefreshUID string `json:"refresh"`
}

type ResetTokenClaims struct {
	UserId string `json:"user_id"`
	UUID   string `json:"id"`
	jwt.RegisteredClaims
}

type ResetTokenCached struct {
	ResetUID string `json:"reset"`
}

type tokenStruct struct{}

func NewToken() *tokenStruct {
	return &tokenStruct{}
}

func (t *tokenStruct) CreateAccess(ctx *context.Context, userId, userRole *string) (*string, *jwt.NumericDate, *Error) {
	tokenString := new(string)
	expired := new(jwt.NumericDate)

	expiredTime := time.Minute * env.NewEnv().JWT_EXPIRED_ACCESS

	tokenUUID := uuid.NewString()
	claims := AccessTokenClaims{
		*userId,
		*userRole,
		tokenUUID,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiredTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return tokenString, expired, &Error{
			Errors: "can't signed the token",
		}
	}

	tokenString = &ss

	cachedJson, err := json.Marshal(AccessTokenCached{
		AccessUID: claims.UUID,
	})
	if err != nil {
		return tokenString, expired, &Error{
			Errors: "can't marshal access token",
		}
	}

	if err := infrastructure.Redis.Set(*ctx, fmt.Sprintf("access-token-%s", *userId), string(cachedJson), expiredTime).Err(); err != nil {
		return tokenString, expired, &Error{
			Errors: "can't cached access token",
		}
	}
	return tokenString, claims.ExpiresAt, &Error{}
}

func (t *tokenStruct) ParseAccess(tokenString *string) (*AccessTokenClaims, *Error) {
	token, err := jwt.ParseWithClaims(*tokenString, &AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return nil, &Error{
			Errors: err.Error(),
		}
	} else if claims, ok := token.Claims.(*AccessTokenClaims); ok {
		return claims, &Error{}
	} else {
		return nil, &Error{
			Errors: "unknown claims type, cannot proceed",
		}
	}
}

func (t *tokenStruct) ValidateAccess(ctx *context.Context, claims *AccessTokenClaims) *Error {
	g := new(errgroup.Group)
	g.Go(func() error {
		cacheJSON, err := infrastructure.Redis.Get(*ctx, fmt.Sprintf("access-token-%s", claims.UserId)).Result()
		if err != nil {
			return fmt.Errorf("token not found")
		}
		cachedTokens := new(AccessTokenCached)
		err = json.Unmarshal([]byte(cacheJSON), cachedTokens)
		var tokenUID string = cachedTokens.AccessUID
		if err != nil || tokenUID != claims.UUID {
			return fmt.Errorf("token not found")
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return &Error{
			Errors: err.Error(),
		}
	}
	return &Error{}
}

func (t *tokenStruct) CreateRefresh(ctx *context.Context, userId, userRole *string) (*string, *jwt.NumericDate, *Error) {
	tokenString := new(string)
	expired := new(jwt.NumericDate)

	expiredTime := time.Minute * env.NewEnv().JWT_EXPIRED_REFRESH

	tokenUUID := uuid.NewString()
	claims := RefreshTokenClaims{
		*userId,
		*userRole,
		tokenUUID,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiredTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return tokenString, expired, &Error{
			Errors: "can't signed the token",
		}
	}

	tokenString = &ss

	cachedJson, err := json.Marshal(RefreshTokenCached{
		RefreshUID: claims.UUID,
	})
	if err != nil {
		return tokenString, expired, &Error{
			Errors: "can't marshal refresh token",
		}
	}

	if err := infrastructure.Redis.Set(*ctx, fmt.Sprintf("refresh-token-%s", *userId), string(cachedJson), expiredTime).Err(); err != nil {
		return tokenString, expired, &Error{
			Errors: "can't cached refresh token",
		}
	}
	return tokenString, claims.ExpiresAt, &Error{}
}

func (t *tokenStruct) ParseRefresh(tokenString *string) (*RefreshTokenClaims, *Error) {
	token, err := jwt.ParseWithClaims(*tokenString, &RefreshTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return nil, &Error{
			Errors: err.Error(),
		}
	} else if claims, ok := token.Claims.(*RefreshTokenClaims); ok {
		return claims, &Error{}
	} else {
		return nil, &Error{
			Errors: fmt.Errorf("unknown claims type, cannot proceed"),
		}
	}
}

func (t *tokenStruct) ValidateRefresh(ctx *context.Context, claims *RefreshTokenClaims) *Error {
	g := new(errgroup.Group)
	g.Go(func() error {
		fmt.Println("cek 12")
		cacheJSON, err := infrastructure.Redis.Get(*ctx, fmt.Sprintf("refresh-token-%s", claims.UserId)).Result()
		if err != nil {
			return fmt.Errorf("token not found")
		}
		cachedTokens := new(RefreshTokenCached)
		err = json.Unmarshal([]byte(cacheJSON), cachedTokens)
		var tokenUID string = cachedTokens.RefreshUID
		if err != nil || tokenUID != claims.UUID {
			return fmt.Errorf("token not found")
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return &Error{
			Errors: err.Error(),
		}
	}
	return &Error{}
}

func (t *tokenStruct) CreateReset(ctx *context.Context, userId *string) (*string, *jwt.NumericDate, *Error) {
	tokenString := new(string)
	expired := new(jwt.NumericDate)

	expiredTime := time.Minute * env.NewEnv().JWT_EXPIRED_RESET

	tokenUUID := uuid.NewString()
	claims := ResetTokenClaims{
		*userId,
		tokenUUID,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiredTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return tokenString, expired, &Error{
			Errors: "can't signed the token",
		}
	}

	tokenString = &ss

	cachedJson, err := json.Marshal(ResetTokenCached{
		ResetUID: claims.UUID,
	})
	if err != nil {
		return tokenString, expired, &Error{
			Errors: "can't marshal reset token",
		}
	}

	if err := infrastructure.Redis.Set(*ctx, fmt.Sprintf("reset-token-%s", *userId), string(cachedJson), expiredTime).Err(); err != nil {
		return tokenString, expired, &Error{
			Errors: "can't cached reset token",
		}
	}
	return tokenString, claims.ExpiresAt, &Error{}
}

func (t *tokenStruct) ParseReset(tokenString *string) (*ResetTokenClaims, *Error) {
	token, err := jwt.ParseWithClaims(*tokenString, &ResetTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return nil, &Error{
			Errors: err.Error(),
		}
	} else if claims, ok := token.Claims.(*ResetTokenClaims); ok {
		return claims, &Error{}
	} else {
		return nil, &Error{
			Errors: "unknown claims type, cannot proceed",
		}
	}
}

func (t *tokenStruct) ValidateReset(ctx *context.Context, claims *ResetTokenClaims) *Error {
	g := new(errgroup.Group)
	g.Go(func() error {
		cacheJSON, err := infrastructure.Redis.Get(*ctx, fmt.Sprintf("reset-token-%s", claims.UserId)).Result()
		if err != nil {
			return fmt.Errorf("token not found")
		}
		cachedTokens := new(ResetTokenCached)
		err = json.Unmarshal([]byte(cacheJSON), cachedTokens)
		var tokenUID string = cachedTokens.ResetUID
		if err != nil || tokenUID != claims.UUID {
			return fmt.Errorf("token not found")
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return &Error{
			Errors: err.Error(),
		}
	}
	return &Error{}
}
