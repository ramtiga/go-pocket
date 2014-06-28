package main

import (
	"fmt"
	"github.com/ramtiga/go-pocket"
	"log"
)

func main() {
	client := pocket.NewClient("CONSUMER_KEY", "ACCESS_TOKEN")

	results, err := client.PocketList()
	if err != nil {
		log.Fatal(err)
	}

	for _, res := range results {
		fmt.Println(res.Resolved_title)
	}
}
