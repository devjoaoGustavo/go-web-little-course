package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	visit := "https://http-methods.appspot.com/Hungary/"
	var resp *http.Response

	req, err := http.NewRequest("GET", visit, nil)
	if err != nil {
		log.Fatal(err)
	}

	query := req.URL.Query()
	query.Add("v", "true")

	req.URL.RawQuery = query.Encode()

	client := http.DefaultClient

	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println(req.URL.String(), resp.Status)
	body, err1 := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body), err1)
}
