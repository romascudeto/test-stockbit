package config

import (
	"soal2/controllers"

	echo "github.com/labstack/echo/v4"
)

func GenRoutes(e *echo.Echo) {
	e.GET("/search", controllers.Search)
	e.GET("/singleByID", controllers.DetailID)
	e.GET("/singleByTitle", controllers.DetailTitle)
}
