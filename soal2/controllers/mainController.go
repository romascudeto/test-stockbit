package controllers

import (
	"net/http"
	"soal2/helper"
	"strings"

	"github.com/labstack/echo/v4"
)

func Search(c echo.Context) error {
	searchword := c.QueryParam("searchword")
	pagination := c.QueryParam("pagination")
	resp := helper.Search(searchword, pagination)
	metaData := map[string]interface{}{
		"totalData": resp.TotalResults,
	}
	respSucc := helper.ResponseSuccess(metaData, resp.Search)
	return c.JSON(http.StatusOK, respSucc)
}

func DetailID(c echo.Context) error {
	idMovie := c.QueryParam("idMovie")
	resp := helper.DetailMovieByID(idMovie)
	respSucc := helper.ResponseSuccess(nil, resp)
	return c.JSON(http.StatusOK, respSucc)
}

func DetailTitle(c echo.Context) error {
	titleMovie := c.QueryParam("titleMovie")
	titleMovie = strings.ReplaceAll(titleMovie, " ", "+")
	resp := helper.DetailMovieByTitle(titleMovie)
	respSucc := helper.ResponseSuccess(nil, resp)
	return c.JSON(http.StatusOK, respSucc)
}
