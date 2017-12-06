package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	visit := "https://http-methods.appspot.com/Avestruxa/Message"
	var resp *http.Response

	reader := strings.NewReader("Aquela string que eu quiser")

	req, err := http.NewRequest("PUT", visit, reader)
	if err != nil {
		log.Fatal(err)
	}
	client := http.DefaultClient

	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println(visit, resp.Status)
	body, err1 := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body), err1)
}
