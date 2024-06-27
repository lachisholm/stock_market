// This will get the stock price for a stock symbol

package main

import (
	"github.com/go-resty/resty/v2" // Import the resty package for making HTTP requests
)

const apiKey = "YOUR_ALPHA_VANTAGE_API_KEY" // Define the API key for Alpha Vantage

// getStockPrice function fetches the stock price for a given symbol
func getStockPrice(symbol string) (string, error) {
	client := resty.New()    // Create a new Resty client
	resp, err := client.R(). // Make a GET request
					SetQueryParams(map[string]string{ // Set query parameters
			"function": "TIME_SERIES_INTRADAY", // API function
			"symbol":   symbol,                 // Stock symbol
			"interval": "1min",                 // Interval for stock data
			"apikey":   apiKey,                 // API key
		}).
		Get("https://www.alphavantage.co/query") // API endpoint

	if err != nil {
		return "", err // Return error if request fails
	}
	return resp.String(), nil // Return response as string if successful
}
