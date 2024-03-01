package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

const addr string = "localhost:8081"

func main() {

	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(addr, nil))

}

func handle(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	text := string(bodyBytes)
	responce := "2 instance: " + text + "\n"

	if _, err := w.Write([]byte(responce)); err != nil {
		log.Fatal(err)
	}
}
