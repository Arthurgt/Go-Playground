package main

import (
	"fmt"
	"net/http"
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

	for i := 0; i < len(links); i++ {
		fmt.Println(<-linkChannel)
	}
}

func checkLink(link string, linkChannel chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " might be down!")
		linkChannel <- "Might be down I think"
		return
	}
	fmt.Println(link + " is up!")
	linkChannel <- "Yep, is up"
}
