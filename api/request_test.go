package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBasicForecastOnGivenCoordinates(t *testing.T) {
	// Coordinates of Firenze Centro Storico
	data, err := GetBasicForecast(43.769599999999997, 11.255800000000001)
	assert.Equal(t, data.Localita.ID, 11526632, "Should be equal")
	assert.Nil(t, err, "Error should be nil")
}

func TestGetBasicForecastOnWrongCoordinates(t *testing.T) {
	_, err := GetBasicForecast(1000000000, 1231312312)

	// assert.Equal(t, data, ForecastBaseResponse{}, "Forecast Base response should be empty")
	assert.NotNil(t, err, "Error should be nil")
}

func TestGetCoordinatesFromSearchTermWithValidParameter(t *testing.T) {

	lat, lon, err := GetCoordinatesFromSearchTerm("firenze")

	assert.NotEqual(t, lat, 0.0, "Lat should not be 0")
	assert.NotEqual(t, lon, 0.0, "Lon should not be 0")
	assert.Nil(t, err, "Error should be nil")

}

func TestGetCoordinatesFromSearchTermWithInvalidParameter(t *testing.T) {

	lat, lon, err := GetCoordinatesFromSearchTerm("wrong parameter")

	assert.Equal(t, lat, 0.0, "Lat should be 0")
	assert.Equal(t, lon, 0.0, "Lon should be 0")
	assert.NotNil(t, err, "Error should not be nil")

}
