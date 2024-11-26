package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func main() {
	// Build the URL with query parameters
	// TODO : Add a parameter to fetch a sepesific kind of quotes
	baseURL := "https://v2.jokeapi.dev/joke/Any"
	params := url.Values{}
	params.Add("format", "txt") // Fetch only funny quotes

	// Construct the final URL
	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// Create an HTTP client with custom transport to skip SSL verification
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // Skip SSL certificate verification
			},
		},
		Timeout: 700 * time.Millisecond,
	}

	// Make the request
	resp, err := client.Get(fullURL)
	if err != nil {
		fmt.Println("You Have No Internet Connection , You Damn Bastard")
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body)) // Display the joke in the terminal
}
