/*
Copyright © 2022 Guglielmo Bartelloni bartelloni.guglielmo@gmail.com

*/
package cmd

import (
	"3bmeteo/api"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search wheater information about a location",
	Run: func(cmd *cobra.Command, args []string) {
		searchTerm := strings.Join(args, " ")
		lat, lon, err := api.GetCoordinatesFromSearchTerm(searchTerm)
		if err != nil {
			panic(err)
		}

		forecastData, err := api.GetBasicForecast(lat, lon)
		if err != nil {
			panic(err)
		}
		for _, day := range forecastData.Localita.PrevisioneGiorno {
			fmt.Println(day.Attendibilita, " ", day.TempoMedio.ProbabilitaPrec, "%")
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
