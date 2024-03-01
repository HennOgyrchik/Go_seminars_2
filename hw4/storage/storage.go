package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"slices"
)

type Storage struct {
	file  *os.File
	users map[int]*User
}

type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Friends []int  `json:"friends"`
}

type UserId struct {
	TargetId int `json:"target_id"`
}

type UpdateAge struct {
	NewAge int `json:"new_age"`
}

type Friends struct {
	SourceId int `json:"source_id"`
	TargetId int `json:"target_id"`
}

func NewStorage(filename string) (*Storage, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return nil, err
	}
	users := make(map[int]*User)

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	if len(data) != 0 {
		err = json.Unmarshal(data, &users)
		if err != nil {
			fmt.Println(1)
			return nil, err
		}
	}

	return &Storage{
		file:  file,
		users: users,
	}, err
}

func (s *Storage) CloseStorage() error {
	data, err := json.Marshal(s.users)
	if err != nil {
		return err
	}
	err = os.WriteFile(s.file.Name(), data, 0666)
	s.file.Close()

	return err
}

func (s *Storage) CreateUser(user User) int {
	s.users[len(s.users)+1] = &user
	return len(s.users)
}

func (s *Storage) MakeFriend(id1, id2 int) {
	s.users[id1].Friends = append(s.users[id1].Friends, id2)
	s.users[id2].Friends = append(s.users[id2].Friends, id1)
}

func (s *Storage) GetName(id int) (string, bool) {
	user, ok := s.users[id]
	if !ok {
		return "", ok
	}
	return user.Name, ok
}

func (s *Storage) DeleteUser(id int) {
	for _, user := range s.users {
		index := slices.Index(user.Friends, id)
		if index == -1 {
			continue
		}
		user.Friends = slices.Delete(user.Friends, index, index+1)
	}
	delete(s.users, id)
}

func (s *Storage) GetIdFriends(id int) []int {
	return s.users[id].Friends
}

func (s *Storage) UpdateAge(id int, age int) {
	s.users[id].Age = age
}
