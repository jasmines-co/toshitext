package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {
	// Set account keys and information
	// Set up Env viriable to hide credentials
	accountSid := "ACCOUNT_SID"
	authToken := "AUTH_TOKEN"
	urlStr := "TWILIO_ENDPOINT_URL"

	// Rand seed
	rand.Seed(time.Now().Unix())

	messages := [7]string{
		"Follow the white rabbit.",
		"Any sufficiently advanced technology is indistinguishable from magic.'",
		"The Times 03/Jan/2009 Chancellor on brink of second bailout for banks.",
		"There is nothing in a caterpillar that tells you it's going to be a butterfly.'",
		"Messages are broadcast on a best effort basis.",
		"Did you get my text?",
		"Bitcoin - A purely peer-to-peer version of electronic cash system.",
	}

	msgData := url.Values{}
	msgData.set("To", "NUMBER_TO")
	msgData.set("From", "NUMBER_FROM")
	msgData.set("Body", quotes[rand.Intn(len(quotes))])
	msgDataReader := *strings.NewReader(msgData.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Conten-Type", "application/x-www.-form-urlencoded")

	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			fmt.Println(data["sid"])
		} else {
			fmt.Println(resp.Status)
		}
	}

	fmt.Println("ToshiText Backend Api vAlpha")
}
