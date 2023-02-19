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
	"time"
)

func main() {
	defer Prompts()

	speech := htgotts.Speech{Folder: "audio", Language: voices.English, Handler: &handlers.Native{}}
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

	if getInput() == "Time" {
		//Check greetTime and Greet accordingly
		greetTime := getTime()
		if greetTime == "morning" {
			err := speech.Speak("Good " + greetTime + " " + getInput())
			if err != nil {
				return

			}
		} else if greetTime == "afternoon" {
			err := speech.Speak("Good " + greetTime + " " + getInput())
			if err != nil {
				return

			}
		} else if greetTime == "evening" {
			err := speech.Speak("Good " + greetTime + " " + getInput())
			if err != nil {
				return

			}
		}
	} else {
		err := speech.Speak("Ron Swanson says: " + ronSwanson())
		log.Println("Ron Swanson says: " + ronSwanson())
		if err != nil {
			return

		}
	}

}

func getTime() string {
	//GET TIME
	//RETURN TIME
	timeNow := time.Now()
	timeString := timeNow.Format("15:04:05")
	if timeString < "12:00:00" {
		return "morning"
	} else if timeString < "18:00:00" {
		return "afternoon"
	} else {
		return "evening"
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

func Prompts() {
	//PROMPTS
	fmt.Println("Thank to the Ron Swanson Quote Generator!")

	//Create back and forth between user and program
	//Get user input
	//Search for quotes
	//Print quotes
	//Ask if user wants to search again
	//If yes, repeat
	//If no, exit

}
