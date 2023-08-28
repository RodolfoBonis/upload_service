package middlewares

import (
	keyGuardian "github.com/RodolfoBonis/go_key_guardian"
	"github.com/labstack/echo/v4"
	"net/http"
)

func KeyGuardian(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		apiKey := c.Request().Header.Get("X-Api-Key")

		principalUser, isValid := keyGuardian.ValidateApiKey(apiKey)

		if !isValid {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "You cannot access this content"})
		}

		c.Set("currentUser", principalUser)

		err := next(c)
		if err != nil {
			return err
		}

		return nil
	}
}
