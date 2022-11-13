package middlewares

import (
	"admin-server/daos/auth"
	"admin-server/vars"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	authHeader = "Auth"
)

func AuthMiddleware() echo.MiddlewareFunc {
	dao := auth.NewDAO()

	invalidError := echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("missing or invalid '%s' header", authHeader))

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get(authHeader)
			if authHeader == "" {
				return invalidError
			}

			// validate token
			user, err := dao.GetUserFromToken(authHeader)
			if err != nil {
				vars.Logger.LogError("middleware:auth.AuthMiddleware()", "failed to get user from token (%s)", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}

			if user == nil {
				return invalidError
			}

			c.Set("user", &user)

			return next(c)
		}
	}
}
