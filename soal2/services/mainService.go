package services

import (
	"encoding/json"
	"os"
	"soal2/models"

	resty "github.com/go-resty/resty/v2"
)

func Search(searchword string, pagination string, respCh chan<- models.Search) {
	client := resty.New()
	var searchMovie models.Search
	resp, _ := client.R().
		EnableTrace().
		Get("http://www.omdbapi.com/?apikey=" + os.Getenv("API_KEY") + "&s=" + searchword + "&page=" + pagination)
	json.Unmarshal(resp.Body(), &searchMovie)
	respCh <- searchMovie
	close(respCh)
}

func DetailMovieByID(idMovie string, respCh chan<- models.Movie) {
	client := resty.New()
	var movie models.Movie
	resp, _ := client.R().
		EnableTrace().
		Get("http://www.omdbapi.com/?apikey=" + os.Getenv("API_KEY") + "&i=" + idMovie)
	json.Unmarshal(resp.Body(), &movie)
	respCh <- movie
	close(respCh)
}

func DetailMovieByTitle(idMovie string, respCh chan<- models.Movie) {
	client := resty.New()
	var movie models.Movie
	resp, _ := client.R().
		EnableTrace().
		Get("http://www.omdbapi.com/?apikey=" + os.Getenv("API_KEY") + "&t=" + idMovie)
	json.Unmarshal(resp.Body(), &movie)
	respCh <- movie
	close(respCh)
}
