package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jonnypillar/what3words/pkg/w3w"
)

func main() {
	apiKey := os.Getenv("W3W_API_KEY")

	c, _ := w3w.New(apiKey)

	res, err := c.GetWords(w3w.Coordinates{
		Lat: 51.432393,
		Lng: -0.348023,
	}, w3w.WordOptions{})

	if err != nil {
		fmt.Println("Error occurred converting coordinates to words", err)

		return
	}

	body, _ := json.MarshalIndent(res, "", "  ")

	fmt.Println("Results")
	fmt.Println(string(body))
}
