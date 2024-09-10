package middleware

import (
	"bitbucket.org/bri_bootcamp/fp-patungan-backend-go/pkg/jwt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")

		if token == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token is missing")
		}

		// remove bearer
		logrus.Print("token cuyyy: ", token)
		if strings.Contains(token, "Bearer ") {
			token = strings.Replace(token, "Bearer ", "", 1)
			logrus.Println("token baru cuyyy: ", token)
		} else {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		// validate
		claims, err := jwt.ValidateToken(token)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		// Set user ID in the context
		c.Set("userID", claims.UserID)
		return next(c)
	}
}
