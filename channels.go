package main

import (
	"fmt"
	"net/http"
	"time" 
)

func makeChannels() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
		"http://stackoverflow.com",
	}

	linkChannel := make(chan string)

	for _, link := range links {
		go checkLink(link, linkChannel)
	}

	for l := range linkChannel {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, linkChannel)
		}(l)
	}
}

func checkLink(link string, linkChannel chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " might be down!")
		linkChannel <- link
		return
	}
	fmt.Println(link + " is up!")
	linkChannel <- link
}
