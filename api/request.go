package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Search the location id using a given string
func SearchLocation(searchTerm string) (float64, float64) {

	url := "https://api.3bmeteo.com/mobilev3/api_localita/ricerca_elastic_search/" + searchTerm + "/it/italia/?format=json2&X-API-KEY=fhrwRdevqwq8r7q9UXTwP6lSX74g34jnQ6756tGo"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	localitaResponse := LocalitaResponse{}

	err = json.Unmarshal(body, &localitaResponse)
	if err != nil {
		fmt.Println(err)
	}

	//Takes the first occurence (needs to be selectable)
	return localitaResponse.Localita[0].Lat, localitaResponse.Localita[0].Lon
}

func GetBasicForecast(lat, lon float64) ForecastBaseResponse {
	forecastBaseResponse := ForecastBaseResponse{}

	url := fmt.Sprintf("https://api.3bmeteo.com/mobilev3/api_previsioni/home_geo/%f/%f/it/?format=json2&X-API-KEY=fhrwRdevqwq8r7q9UXTwP6lSX74g34jnQ6756tGo", lat, lon)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(body, &forecastBaseResponse)
	if err != nil {
		fmt.Println(err)
	}
	return forecastBaseResponse
}
