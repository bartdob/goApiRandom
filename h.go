package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
			"apiKey": "*************************",
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
	fmt.Printf("%#v\n", anwser)
	fmt.Printf("------------------------------")
	fmt.Printf("%#v\n", anwser["result"])

	// for k, v := range anwser {
	// 	fmt.Printf("key is %v and val %v\n Type: %T\n", k, v, v)
	// }
	fmt.Printf("****************************************************************************")
	u := JsonResponse{}

	json.Unmarshal([]byte(body), &u)

	// for i := 0; i < 1; {
	// 	fmt.Printf("Code is: %s", u.Result.Random.Data[i])
	// }

	}

}
