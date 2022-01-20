# goApiRandom
https://api.random.org/json-rpc/4/basic api to get random numbers, base on your json data:

# Configuration
```json

{
	"jsonrpc": "2.0",
	"method": "generateIntegers",
	"params": {
		"apiKey": #"your passord to random.org",
		"n": 6, # how many numbers would you like to get
		"min": 1, #range
		"max": 100, #range
		"replacement": true
	},
	"id": 42
}
```

