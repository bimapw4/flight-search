package common

import (
	"context"
	"flight-api/internal/entity"
)

func SetUserCtx(c context.Context, claims *entity.Claim) context.Context {
	return context.WithValue(c, "user_ctx", claims)
}

func GetUserCtx(c context.Context) *entity.Claim {
	return c.Value("user_ctx").(*entity.Claim)
}

func SetTokenCtx(c context.Context, token string) context.Context {
	return context.WithValue(c, "token", token)
}

func GetTokenCtx(c context.Context) string {
	return c.Value("token").(string)
}
