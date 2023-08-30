// This is based on the the course work from Boot.dev on Pointers in Go
// check out Boot.dev for some great instructional learning

package main

import (
	"fmt"
	"strings"
)

func removeProfanity(message *string) {

	msg := *message
	badWords := []string{
		"dang",
		"shoot",
		"heck",
	}

	for _, word := range badWords {
		censored := ""
		for range word {
			censored += "*"
		}
		msg = strings.ReplaceAll(msg, word, censored)
	}
	*message = msg
}


func test(messages []string) {
	for _, message := range messages {
		removeProfanity(&message)
		fmt.Println(message)
	}
}

func main() {
	messages1 := []string{
		"",
		"dang robots",
		"dang them to heck",
	}

	messages2 := []string{
		"well shoot",
		"Allan is going straight to heck",
		"dang... that's a tough break",
	}

	test(messages1)
	test(messages2)
}
