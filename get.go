package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	sites := []string{"https://golang.org", "http://golang.org", "https://golang.org/foo"}
	var resp *http.Response
	var err error

	for _, site := range sites {
		resp, err = http.Get(site)
		if err != nil {
			log.Fatal(err)
		}

		if resp.StatusCode == 404 {
			fmt.Println(site, resp.Status)
		}
	}
}
