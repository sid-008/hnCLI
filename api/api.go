package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func GetTopItems() {
	resp, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	sb := strings.Split(string(body), ",")
	log.Printf("%s", sb)
	log.Println("Size is:", len(sb))
	// return sb
}

func GetTopItemsLinks() {
	// items := make(map[int]string)

}
