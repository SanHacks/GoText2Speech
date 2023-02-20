package main

import (
	"github.com/hegedustibor/htgo-tts"
	"time"
)

// Get Time of Day and return it as a string
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

func tellTime(speech htgotts.Speech) bool {
	//Check greetTime and Greet accordingly
	greetTime := getTime()
	if greetTime == "morning" {
		err := speech.Speak("Good " + greetTime + " " + getInput())
		if err != nil {
			return true

		}

	} else if greetTime == "afternoon" {
		err := speech.Speak("Good " + greetTime + " " + getInput())
		if err != nil {
			return true

		}
	} else if greetTime == "evening" {
		err := speech.Speak("Good " + greetTime + " " + getInput())
		if err != nil {
			return true

		}
	}
	return false
}
