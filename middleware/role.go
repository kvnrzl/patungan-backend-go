package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func RoleMiddleware(allowedRoles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userID := c.Get("userID").(string)

			// Assume we fetch user role from the database using userID
			userRole := getUserRole(userID) // Function to get role from DB

			// Check if user's role is allowed
			for _, role := range allowedRoles {
				if userRole == role {
					return next(c)
				}
			}
			return echo.NewHTTPError(http.StatusForbidden, "You don't have permission to access this resource")
		}
	}
}

func getUserRole(userID string) string {
	// Dummy function, replace this with actual DB lookup
	if userID == "1" {
		return "admin"
	}
	return "user"
}
