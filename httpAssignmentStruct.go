package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"time"
)

const GetMoviesURL = "https://jsonmock.hackerrank.com/api/movies/search/?Title=substr"

type MoviesData struct {
	Data []Movies `json:"data"`
}

type Movies struct {
	Title  string
	Year   float64
	ImdbID string
}

func main() {
	var title = "life"
	var pageNumber = "2"
	var inputYear = "2008"
	movieTitles := getMovieTitles(title, pageNumber, inputYear)
	fmt.Println("\nMovies in ascending order by name:")
	for _, v := range movieTitles {
		fmt.Println(v)
	}

	movieTitles = getMovieTitles(title, "", "")
	fmt.Println("\nMovies in ascending order by name:")
	for _, v := range movieTitles {
		fmt.Println(v)
	}
}

func getMovieTitles(substr string, pageNumber string, year string) []string {
	moviesJSON := getMoviesJSON(substr, pageNumber)
	return parseMovies(moviesJSON, year)
}

func getMoviesJSON(substr string, pageNumber string) []byte {
	client := http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, formRequestURL(substr, pageNumber), nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	defer res.Body.Close()

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	fmt.Println("\n\nResponse--->", string(body))
	return body
}

func formRequestURL(substr string, pageNumber string) string {
	u, err := url.Parse(GetMoviesURL)
	if err != nil {
		log.Fatal(err)
	}

	q := u.Query()
	q.Set("Title", substr)
	if pageNumber == "" {
		pageNumber = "1"
	}
	q.Add("page", pageNumber)
	u.RawQuery = q.Encode()
	return u.String()
}

func parseMovies(moviesJSON []byte, inputYear string) []string {
	var movieTitles []string

	inputYearF, _ := strconv.ParseFloat(inputYear, 64)
	// var parsedResult = new(MoviesData)
	parsedResult := MoviesData{}
	jsonErr := json.Unmarshal(moviesJSON, &parsedResult)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	data := parsedResult.Data

	for _, movie := range data {
		year := movie.Year
		if inputYear == "" || inputYearF == year {
			title := movie.Title
			movieTitles = append(movieTitles, title)
		}
	}

	sort.Strings(movieTitles)
	return movieTitles
}
