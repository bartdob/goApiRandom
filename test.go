package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/CDFN/rangomorg"
)

const (
	randomApiKey = "0a48fc40-5066-4794-8dc0-036e41d09c63"
)

func main() {
	random := rangomorg.New(randomApiKey)
	result, err := random.GenerateSignedStrings(5, 10, "rangom", map[string]interface{}{
		"userData":    "YourUserData", // These options are optional
		"replacement": false,           // see https://api.random.org/json-rpc/2 for more
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	jsonBytes, _ := json.MarshalIndent(result, "", "    ")
	fmt.Println("Random: ")
	fmt.Println(string(jsonBytes))                      // Display result in json form
	fmt.Println("Requested data: ", result.Random.Data) // Display requested data
	fmt.Println("Signature: ", result.Signature)        // In case of signed api, display signature
}
