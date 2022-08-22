/*
Copyright © 2022 Guglielmo Bartelloni bartelloni.guglielmo@gmail.com

*/
package cmd

import (
	"3bmeteo/api"
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var (
	mainColor = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#FCE2DB", Dark: "#FCE2DB"}).
			Render
	secondColor = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#FF8FB1", Dark: "#FF8FB1"}).
			Render
	thirdColor = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#B270A2", Dark: "#B270A2"}).
			Render
	fourthColor = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#7A4495", Dark: "#7A4495"}).
			Render
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

		fmt.Println(mainColor(fmt.Sprintf("%s, %s", forecastData.Localita.Localita, forecastData.Localita.Prov)))
		fmt.Println(secondColor(fmt.Sprintf("Accuratezza: %s", forecastData.Localita.PrevisioneGiorno[0].Attendibilita)))
		fmt.Println(thirdColor(fmt.Sprintf("Condizioni attuali: %s", forecastData.Localita.PrevisioneGiorno[0].PrevisioneOraria[0].DescBreve)))
		fmt.Println(fourthColor(fmt.Sprintf("Temperature: %dC, %dC", forecastData.Localita.PrevisioneGiorno[0].TempoMedio.TMin, forecastData.Localita.PrevisioneGiorno[0].TempoMedio.TMax)))
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
