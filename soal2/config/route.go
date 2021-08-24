package config

import (
	"net/http"

	helper "soal2/helper"

	echo "github.com/labstack/echo/v4"
)

func GenRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/search", func(c echo.Context) error {
		searchword := c.QueryParam("searchword")
		pagination := c.QueryParam("pagination")
		resp := helper.Search(searchword, pagination)
		metaData := map[string]interface{}{
			"totalData": resp.TotalResults,
		}
		respSucc := helper.ResponseSuccess(metaData, resp.Search)
		return c.JSON(http.StatusOK, respSucc)
	})
	e.GET("/single", func(c echo.Context) error {
		idMovie := c.QueryParam("idMovie")
		resp := helper.DetailMovie(idMovie)
		respSucc := helper.ResponseSuccess(nil, resp)
		return c.JSON(http.StatusOK, respSucc)
	})
}
