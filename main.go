package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getWeather(city string, country string) (string, error) {
	url := fmt.Sprintf("https://www.timeanddate.com/weather/%s/%s", country, city)

	// make an HTTP request
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to make HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Check for HTTP response status
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get a valid response: %s", resp.Status)
	}

	// parse the HTML page
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to parse HTML: %w", err)
	}

	// extract the weather
	weather := doc.Find(".h2").Text()
	if weather == "" {
		return "", fmt.Errorf("could not find weather information for %s, %s", city, country)
	}
	weather = weather[0:2] + "°C" // get only the temperature and add the unit

	return weather, nil
}

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Please provide a country and city as arguments.")
	}

	// parse user input and format it to be used in the URL
	city := strings.ReplaceAll(os.Args[1], " ", "")
	country := strings.ReplaceAll(os.Args[2], " ", "")

	// get the weather and handle any errors
	weather, err := getWeather(city, country)
	if err != nil {
		log.Fatal(err)
	}

	// print the weather
	fmt.Printf("The weather in %s, %s is %s.\n", city, country, weather)
}
