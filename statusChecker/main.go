package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://amazon.com",
		"http://facebook.com",
		"http://golang.org",
		"http://google.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Link: ", link, " is down!")
		c <- "Might be down!"
		return
	}
	fmt.Println("Link: ", link, " is Up!")
	c <- "it's Up!"
}
