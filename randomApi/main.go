package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
)

type JsonResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Random struct {
			Data           []int  `json:"data"`
			CompletionTime string `json:"completionTime"`
		} `json:"random"`
		BitsUsed      int `json:"bitsUsed"`
		BitsLeft      int `json:"bitsLeft"`
		RequestsLeft  int `json:"requestsLeft"`
		AdvisoryDelay int `json:"advisoryDelay"`
	} `json:"result"`
	ID int `json:"id"`
}

func main() {
	httpposturl := "https://api.random.org/json-rpc/4/invoke"
	fmt.Println("HTTP JSON POST URL:", httpposturl)

	var jsonData = []byte(`{
		"jsonrpc": "2.0",
		"method": "generateIntegers",
		"params": {
			"apiKey": "password",
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

	checkVaild := json.Valid(jsonData)

	if checkVaild {
		fmt.Println("Json valid")
	}

	var anwser map[string]interface{}
	json.Unmarshal([]byte(body), &anwser)

	u := JsonResponse{}

	json.Unmarshal([]byte(body), &u)

	var suma float64

	for i := range u.Result.Random.Data {
		fmt.Println(i, u.Result.Random.Data[i])
		suma = suma + float64(u.Result.Random.Data[i])
	}

	fmt.Println("suma: ", suma)

	fmt.Println("standard deviation")

	var mean, sd float64

	mean = suma / float64(len(u.Result.Random.Data))

	for j := range u.Result.Random.Data {
		sd += math.Pow(float64(u.Result.Random.Data[j])-mean, 2)
	}

	sd = math.Sqrt(sd / float64(len(u.Result.Random.Data)))

	fmt.Println("The Standard Deviation is : ", sd, "n:", float64(len(u.Result.Random.Data)))

}
