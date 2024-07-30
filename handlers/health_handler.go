package handlers

import (
	"net/http"

	"github.com/RodolfoBonis/upload_service/helper"
	"github.com/RodolfoBonis/upload_service/models"
	"github.com/labstack/echo/v4"
)

type HealthHandler interface {
	GetHealth(c echo.Context) error
}

type health struct{}

func NewHealthHandler() HealthHandler {
	return &health{}
}

func (_ health) GetHealth(c echo.Context) error {

	ip, err := helper.GetLocalIP()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, models.IpModel{
		Address: ip.String(),
		Health:  true,
	})
}
