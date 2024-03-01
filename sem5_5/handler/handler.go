package handler

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	text := "message from test handler"
	if _, err := w.Write([]byte(text)); err != nil {
		log.Fatal(err)
	}
}
