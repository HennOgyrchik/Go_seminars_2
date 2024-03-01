package service

import (
	"encoding/json"
	"fmt"
	"hw4/config"
	"hw4/storage"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Service struct {
	Storage *storage.Storage
	Config  config.Config
}

func NewService() (*Service, error) {
	conf, err := config.NewConfig()
	if err != nil {
		return &Service{}, err
	}

	store, err := storage.NewStorage(conf.GetFilename())

	return &Service{Storage: store, Config: conf}, err
}
func (s *Service) Close() {
	s.Storage.CloseStorage()
}

func (s *Service) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var u storage.User
		if err := json.Unmarshal(content, &u); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		id := s.Storage.CreateUser(u)

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf("User was created %d", id)))
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}

func (s *Service) MakeFriends(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var f storage.Friends
		if err := json.Unmarshal(content, &f); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		u1, ok := s.Storage.GetName(f.SourceId)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Пользователь не найден"))
			return
		}
		u2, ok := s.Storage.GetName(f.TargetId)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Пользователь не найден"))
			return
		}

		s.Storage.MakeFriend(f.SourceId, f.TargetId)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("%s и %s теперь друзья", u1, u2)))
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}

func (s *Service) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var u storage.UserId
		if err := json.Unmarshal(content, &u); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		username, ok := s.Storage.GetName(u.TargetId)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Пользователь не найден"))
			return
		}

		s.Storage.DeleteUser(u.TargetId)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("%s удален", username)))
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}

func (s *Service) GetFriends(w http.ResponseWriter, r *http.Request) {
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

		idFriends := s.Storage.GetIdFriends(id)
		result := make([]string, 0, len(idFriends))

		for i, idFriend := range idFriends {
			username, ok := s.Storage.GetName(idFriend)
			if !ok {
				fmt.Println(i, ok)
				continue
			}

			result = append(result, username)
		}

		resp := strings.Join(result, ", ")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(resp))
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}

func (s *Service) UpdateAge(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var u storage.UpdateAge
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

		username, ok := s.Storage.GetName(id)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Пользователь не найден"))
			return
		}
		s.Storage.UpdateAge(id, u.NewAge)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Возраст %s успешно обновлен на %d", username, u.NewAge)))
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}
