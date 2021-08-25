package services

import (
	"encoding/json"
	"os"
	"soal2/models"
	"soal2/repositories"

	resty "github.com/go-resty/resty/v2"
)

func Search(searchword string, pagination string, respCh chan<- models.Search) {
	client := resty.New()
	var searchMovie models.Search
	var url = "http://www.omdbapi.com/?apikey=" + os.Getenv("API_KEY") + "&s=" + searchword + "&page=" + pagination
	resp, _ := client.R().
		EnableTrace().
		Get(url)
	json.Unmarshal(resp.Body(), &searchMovie)
	respCh <- searchMovie
	CreateLogs(url, "Search", resp)
	close(respCh)
}

func DetailMovieByID(idMovie string, respCh chan<- models.Movie) {
	client := resty.New()
	var movie models.Movie
	var url = "http://www.omdbapi.com/?apikey=" + os.Getenv("API_KEY") + "&i=" + idMovie
	resp, _ := client.R().
		EnableTrace().
		Get(url)
	json.Unmarshal(resp.Body(), &movie)
	respCh <- movie
	CreateLogs(url, "Search By ID", resp)
	close(respCh)
}

func DetailMovieByTitle(idMovie string, respCh chan<- models.Movie) {
	client := resty.New()
	var movie models.Movie
	var url = "http://www.omdbapi.com/?apikey=" + os.Getenv("API_KEY") + "&t=" + idMovie
	resp, _ := client.R().
		EnableTrace().
		Get(url)
	json.Unmarshal(resp.Body(), &movie)
	respCh <- movie
	CreateLogs(url, "Search By Title", resp)
	close(respCh)
}

func CreateLogs(url string, activity string, resp *resty.Response) {
	var dataLog models.Log
	dataLog.Activity = activity
	dataLog.URL = url
	dataLog.Request = ""
	dataLog.Response = resp.Body()
	dataLog.ResponseStatus = int32(resp.StatusCode())
	go repositories.CreateLog(&dataLog)
}
