package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	sites := []string{"https://http-methods.appspot.com/Avestruxa/Message"}

	var resp *http.Response

	for _, site := range sites {
		req, err := http.NewRequest("GET", site, nil)
		if err != nil {
			log.Fatal(err)
		}
		client := http.DefaultClient

		resp, err = client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		fmt.Println(site, resp.Status)
		body, err1 := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body), err1)

	}
}
