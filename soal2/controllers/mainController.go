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
	if len(resp.Search) == 0 {
		respErr := helper.ResponseError("Data not found !")
		return c.JSON(http.StatusOK, respErr)
	}
	respSucc := helper.ResponseSuccess("Data fetch successfully", resp.Search)
	return c.JSON(http.StatusOK, respSucc)
}

func DetailID(c echo.Context) error {
	idMovie := c.QueryParam("idMovie")
	resp := helper.DetailMovieByID(idMovie)
	respSucc := helper.ResponseSuccess("Data "+idMovie+" fetch successfully", resp)
	return c.JSON(http.StatusOK, respSucc)
}

func DetailTitle(c echo.Context) error {
	titleMovieParam := c.QueryParam("titleMovie")
	titleMovie := strings.ReplaceAll(titleMovieParam, " ", "+")
	resp := helper.DetailMovieByTitle(titleMovie)
	respSucc := helper.ResponseSuccess("Data "+titleMovieParam+" fetch successfully", resp)
	return c.JSON(http.StatusOK, respSucc)
}
