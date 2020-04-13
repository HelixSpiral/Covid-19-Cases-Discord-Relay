package main

import (
	"bytes"                               // Needed for bytes to send for Discord Webhook
	"encoding/json"                       // Needed for JSON encoding
	"fmt"                                 // Needed for printing
	"github.com/aws/aws-lambda-go/lambda" // Needed for Lambda handler
	"io/ioutil"                           // Needed to read from the web response
	"net/http"                            // Needed for http calls
	"os"                                  // Needed to read env variables
	"strconv"                             // Needed for the int formatting
)

const DEBUG = false

var DiscordWebHook = os.Getenv("DISCORD_WEBHOOK")

type COVID19 struct {
	Total     int64 `json:"cases"`
	Deaths    int64 `json:"deaths"`
	Recovered int64 `json:"recovered"`
}

func lambdaHandler() {
	var cases COVID19

	resp, err := http.Get("https://coronavirus-19-api.herokuapp.com/all")
	handleErr(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	handleErr(err)

	err = json.Unmarshal(body, &cases)
	handleErr(err)

	if DEBUG {
		fmt.Println(cases)
	}

	DiscordMessage := BuildMessage(fmt.Sprintf("```\\nCases: %13s\\nDeaths: %10s\\nRecovered: %7s\\n```", Format(cases.Total), Format(cases.Deaths), Format(cases.Recovered)))
	PostWebhook(DiscordWebHook, DiscordMessage)
}

func main() {
	//lambdaHandler()
	lambda.Start(lambdaHandler)
}

// Format ints with commas.
func Format(n int64) string {
	in := strconv.FormatInt(n, 10)
	numOfDigits := len(in)
	if n < 0 {
		numOfDigits-- // First character is the - sign (not a digit)
	}
	numOfCommas := (numOfDigits - 1) / 3

	out := make([]byte, len(in)+numOfCommas)
	if n < 0 {
		in, out[0] = in[1:], '-'
	}

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = ','
		}
	}
}

func BuildMessage(message string) []byte {
	return []byte(fmt.Sprintf("{\"content\": \"%s\"}", message))
}

func PostWebhook(url string, message []byte) {
	httpClient := http.DefaultClient

	if DEBUG {
		fmt.Println(message)
		fmt.Println(string(message))
	}

	response, err := httpClient.Post(url, "application/json", bytes.NewReader(message))

	if DEBUG {
		fmt.Println(httpClient)
		fmt.Println(response, err)
	}
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
