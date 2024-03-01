package main

import (
	"fmt"
	"hw4/service"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	service, err := service.NewService()
	if err != nil {
		fmt.Println(err)
		return
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("exit")
				service.Close()
				os.Exit(0)
			default:

			}
		}
	}()

	mux := http.NewServeMux()

	mux.HandleFunc("/create", service.Create)
	mux.HandleFunc("/make_friends", service.MakeFriends)
	mux.HandleFunc("/user", service.DeleteUser)
	mux.HandleFunc("/friends/", service.GetFriends)
	mux.HandleFunc("/", service.UpdateAge)

	http.ListenAndServe(service.Config.GetAddress(), mux)
}
