package middleware

import (
	"flight-api/internal/business"
	"flight-api/internal/common"
	"flight-api/internal/response"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Authentication struct {
	business business.Business
	jwt      common.JwtCode
}

func NewAuthentication(business business.Business) *Authentication {
	return &Authentication{
		business: business,
		jwt:      common.NewJwt(),
	}
}

func (m *Authentication) Authentication(c *fiber.Ctx) error {
	authorizationRaw := c.GetReqHeaders()["Authorization"]
	if len(authorizationRaw) < 1 {
		return response.NewResponse("authentication").
			Errors("authentication failled", "unauthenticated").
			JSON(c, fiber.StatusUnauthorized)
	}

	authorization := strings.TrimSpace(authorizationRaw[0])

	splitAuthorization := strings.Split(authorization, " ")
	if len(splitAuthorization) != 2 {
		return response.NewResponse("authentication").
			Errors("authentication failled", "unauthenticated").
			JSON(c, fiber.StatusUnauthorized)
	}

	if splitAuthorization[0] != "Bearer" {
		return response.NewResponse("authentication").
			Errors("authentication failled", "unauthenticated").
			JSON(c, fiber.StatusUnauthorized)
	}

	claim, err := m.jwt.DecodeAccessToken(splitAuthorization[1])
	if err != nil {
		return response.NewResponse("authentication").
			Errors("authentication failled", "unauthenticated").
			JSON(c, fiber.StatusUnauthorized)
	}

	c.SetUserContext(common.SetUserCtx(c.UserContext(), claim))
	c.SetUserContext(common.SetTokenCtx(c.UserContext(), splitAuthorization[1]))

	return c.Next()
}
