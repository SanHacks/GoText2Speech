package main

import (
	"encoding/json"
	"fmt"
	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/handlers"
	"github.com/hegedustibor/htgo-tts/voices"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	speech := htgotts.Speech{Folder: "audio", Language: voices.English, Handler: &handlers.Native{}}

	//loop := true

	var ron = getInput()

	var ronSwanson = func() string {

		url := "https://ron-swanson-quotes.herokuapp.com/v2/quotes"
		method := "GET"

		client := &http.Client{}
		req, err := http.NewRequest(method, url, nil)

		if err != nil {
			fmt.Println(err)
			return ""
		}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(res.Body)

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		var quotes []string
		err = json.Unmarshal(body, &quotes)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		log.Println(quotes[0])
		return quotes[0]
	}

	err := speech.Speak("You Said: " + ron + " Ron Swanson says: " + ronSwanson())
	log.Println("You Said: " + ron + " Ron Swanson says: " + ronSwanson())
	if err != nil {
		return
	}

}

func ronSwanson() string {

	url := "https://ron-swanson-quotes.herokuapp.com/v2/quotes"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(string(body))
	return string(body)
}

type RonSwan struct {
	Quote string `json:"quote"`
}

/*func displayRonSwanson() any {
	var ron RonSwan
	err := json.Unmarshal([]byte(ronSwanson()), &ron)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ron.Quote)
	var hisQuote = "Ron Swanson: " + ron.Quote

	return hisQuote
}

*/

func getInput() string {
	//TAKE CLI INPUT
	var input string
	fmt.Println("Enter a word or phrase to search for: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		return ""
	}
	return input
}
