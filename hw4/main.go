package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type service struct {
	store map[int]*User
}

type UserId struct {
	TargetId int `json:"target_id"`
}
type Friends struct {
	SourceId int `json:"source_id"`
	TargetId int `json:"target_id"`
}
type User struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Friends []string `json:"friends"`
}
type UpdateAge struct {
	NewAge int `json:"new_age"`
}

func main() {
	mux := http.NewServeMux()
	srv := service{make(map[int]*User)}
	mux.HandleFunc("/create", srv.Create)
	mux.HandleFunc("/make_friends", srv.MakeFriends)
	mux.HandleFunc("/user", srv.DeleteUser)
	mux.HandleFunc("/friends/", srv.GetFriends)
	mux.HandleFunc("/", srv.UpdateAge)

	http.ListenAndServe("localhost:8080", mux)
}

func (s *service) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var u User
		if err := json.Unmarshal(content, &u); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		id := len(s.store) + 1
		s.store[id] = &u

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf("User was created %d", id)))
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}

func (s *service) MakeFriends(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var f Friends
		if err := json.Unmarshal(content, &f); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		s.store[f.SourceId].Friends = append(s.store[f.SourceId].Friends, s.store[f.TargetId].Name)
		s.store[f.TargetId].Friends = append(s.store[f.TargetId].Friends, s.store[f.SourceId].Name)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("%s и %s теперь друзья", s.store[f.SourceId].Name, s.store[f.TargetId].Name)))
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}

func (s *service) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var u UserId
		if err := json.Unmarshal(content, &u); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		username := s.store[u.TargetId].Name
		delete(s.store, u.TargetId)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("%s удален", username)))
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}

func (s *service) GetFriends(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		defer r.Body.Close()

		splittedString := strings.Split(r.RequestURI, "/")
		if len(splittedString) < 3 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(splittedString[2])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		friends := s.store[id].Friends

		resp := strings.Join(friends, ", ")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(resp))
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}

func (s *service) UpdateAge(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var u UpdateAge
		if err := json.Unmarshal(content, &u); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		splittedString := strings.Split(r.RequestURI, "/")
		if len(splittedString) < 2 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(splittedString[1])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		s.store[id].Age = u.NewAge

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Возраст %s успешно обновлен на %d", s.store[id].Name, s.store[id].Age)))
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}
