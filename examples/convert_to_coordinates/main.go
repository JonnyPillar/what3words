package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jonnypillar/what3words/pkg/w3w"
)

func main() {
	apiKey := os.Getenv("W3W_API_KEY")

	c := w3w.New(apiKey)

	res, err := c.GetCoordinates(&w3w.Words{"filled", "count"}, &w3w.Options{})

	if err != nil {
		fmt.Println("Error occurred converting words to coordinates", err)

		return
	}

	body, _ := json.MarshalIndent(res, "", "  ")

	fmt.Println("Results")
	fmt.Println(string(body))
}
