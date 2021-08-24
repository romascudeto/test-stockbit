package controllers

import (
	"net/http"
	"soal2/helper"
	"soal2/models"
	"soal2/services"
	"strings"

	"github.com/labstack/echo/v4"
)

func Search(c echo.Context) error {
	var respCh = make(chan models.Search)
	searchword := c.QueryParam("searchword")
	pagination := c.QueryParam("pagination")
	go services.Search(searchword, pagination, respCh)

	var resp models.Search
	for respFor := range respCh {
		resp = respFor
	}

	if len(resp.Search) == 0 {
		respErr := helper.ResponseError("Data not found !")
		return c.JSON(http.StatusOK, respErr)
	}
	respSucc := helper.ResponseSuccess("Data fetch successfully", resp.Search)
	return c.JSON(http.StatusOK, respSucc)
}

func DetailID(c echo.Context) error {
	var respCh = make(chan models.Movie)
	idMovie := c.QueryParam("idMovie")
	go services.DetailMovieByID(idMovie, respCh)

	var resp models.Movie
	for respFor := range respCh {
		resp = respFor
	}

	if resp.Actors == "" {
		respErr := helper.ResponseError("Data not found !")
		return c.JSON(http.StatusOK, respErr)
	}
	respSucc := helper.ResponseSuccess("Data "+idMovie+" fetch successfully", resp)
	return c.JSON(http.StatusOK, respSucc)
}

func DetailTitle(c echo.Context) error {
	var respCh = make(chan models.Movie)
	titleMovieParam := c.QueryParam("titleMovie")
	titleMovie := strings.ReplaceAll(titleMovieParam, " ", "+")
	go services.DetailMovieByTitle(titleMovie, respCh)

	var resp models.Movie
	for respFor := range respCh {
		resp = respFor
	}
	if resp.Actors == "" {
		respErr := helper.ResponseError("Data not found !")
		return c.JSON(http.StatusOK, respErr)
	}
	respSucc := helper.ResponseSuccess("Data "+titleMovieParam+" fetch successfully", resp)
	return c.JSON(http.StatusOK, respSucc)
}
