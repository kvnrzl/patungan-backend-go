package middleware

import (
	"bitbucket.org/bri_bootcamp/patungan-backend-go/models"
	"bitbucket.org/bri_bootcamp/patungan-backend-go/pkg/db"
	"bitbucket.org/bri_bootcamp/patungan-backend-go/pkg/jwt"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")

		if token == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, models.BaseResponse[string]{
				Status:  "failed",
				Message: "Unauthorized - token is missing",
			})
		}

		// remove bearer
		if strings.Contains(token, "Bearer ") {
			token = strings.Replace(token, "Bearer ", "", 1)
		} else {
			return echo.NewHTTPError(http.StatusUnauthorized, models.BaseResponse[string]{
				Status:  "failed",
				Message: "Unauthorized - invalid token format",
			})
		}

		// validate
		claims, err := jwt.ValidateToken(token)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, models.BaseResponse[string]{
				Status:  "failed",
				Message: "Unauthorized - invalid token",
			})
		}

		// check to redis
		redisClient := db.ConnectToRedis()
		if redisClient.Get(context.Background(), claims.Email).Val() == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, models.BaseResponse[string]{
				Status:  "failed",
				Message: "Unauthorized - invalid token",
			})
		}

		// Set user ID in the context
		c.Set("userID", claims.UserID)
		c.Set("name", claims.Name)
		c.Set("email", claims.Email)
		c.Set("phone", claims.Phone)

		return next(c)
	}
}
