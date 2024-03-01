package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const proxyAddr string = "localhost:9000"

var (
	counter int    = 0
	first   string = "http://localhost:8080"
	second  string = "http://localhost:8081"
)

func main() {

	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(proxyAddr, nil))

}

func handle(w http.ResponseWriter, r *http.Request) {
	textBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	text := string(textBytes)

	if counter == 0 {
		resp, err := http.Post(first, "text/plain", bytes.NewBuffer([]byte(text)))
		if err != nil {
			log.Fatal(err)
		}
		counter++

		textBytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		fmt.Println(string(textBytes))
		return
	}

	resp, err := http.Post(second, "text/plain", bytes.NewBuffer([]byte(text)))
	if err != nil {
		log.Fatal(err)
	}
	textBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println(string(textBytes))
	counter--

}
