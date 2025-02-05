package middleware

import (
	"net/http"

	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/pkg/config"
	echoJwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// JWTMiddleware applies JWT authentication to protected routes
func EchoJWTMiddleware() echo.MiddlewareFunc {
	return echoJwt.WithConfig(echoJwt.Config{
		SigningKey:    []byte(config.GetConfig().JwtSecret),
		SigningMethod: "HS256",
		ErrorHandler: func(c echo.Context, err error) error {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or missing token")
		},
	})
}
