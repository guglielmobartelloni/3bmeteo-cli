package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Search the location id using a given string
func GetCoordinatesFromSearchTerm(searchTerm string) (float64, float64, error) {

	url := "https://api.3bmeteo.com/mobilev3/api_localita/ricerca_elastic_search/" + searchTerm + "/it/italia/?format=json2&X-API-KEY=fhrwRdevqwq8r7q9UXTwP6lSX74g34jnQ6756tGo"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return 0, 0, err
	}
	res, err := client.Do(req)
	if err != nil {
		return 0, 0, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, 0, err
	}

	localitaResponse := LocalitaResponse{}
	if strings.Contains(string(body), "error") {
		return 0, 0, errors.New("The response from the api returned an error")
	}

	err = json.Unmarshal(body, &localitaResponse)
	if err != nil {
		return 0, 0, err
	}

	//Takes the first occurence (needs to be selectable)
	return localitaResponse.Localita[0].Lat, localitaResponse.Localita[0].Lon, nil
}

func GetBasicForecast(lat, lon float64) (ForecastBaseResponse, error) {
	forecastBaseResponse := ForecastBaseResponse{}

	url := fmt.Sprintf("https://api.3bmeteo.com/mobilev3/api_previsioni/home_geo/%f/%f/it/?format=json2&X-API-KEY=fhrwRdevqwq8r7q9UXTwP6lSX74g34jnQ6756tGo", lat, lon)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return forecastBaseResponse, err
	}
	res, err := client.Do(req)
	if err != nil {
		return forecastBaseResponse, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return forecastBaseResponse, err
	}
	if strings.Contains(string(body), "error") {
		return forecastBaseResponse, errors.New("The response from the api returned an error")
	}

	err = json.Unmarshal(body, &forecastBaseResponse)
	if err != nil {
		return forecastBaseResponse, err
	}

	return forecastBaseResponse, nil
}
