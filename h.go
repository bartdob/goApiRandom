package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	httpposturl := "https://api.random.org/json-rpc/4/invoke"
	fmt.Println("HTTP JSON POST URL:", httpposturl)

	var jsonData = []byte(`{
		"jsonrpc": "2.0",
		"method": "generateIntegers",
		"params": {
			"apiKey": "**********************",
			"n": 6,
			"min": 1,
			"max": 100,
			"replacement": true
		},
		"id": 42
	}`)
	request, error := http.NewRequest("POST", httpposturl, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	fmt.Println("response Status:...", response.Status)
	// fmt.Println("response Headers:...", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body[45]))
	fmt.Println("response Body:", string(body))

	json.NewDecoder(response.Body).Decode(&res)

	fmt.Println(response["form"])

}
