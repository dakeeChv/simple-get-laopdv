package router

import (
	"lao-pdv/src/controller"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/province", controller.GetProvince)
	e.GET("/district", controller.GetDistrict)
	e.GET("/village", controller.GetVillage)
	return e
}
