package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Search the location id using a given string
func SearchId(searchTerm string) int {

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
	return localitaResponse.Localita[0].ID
}
