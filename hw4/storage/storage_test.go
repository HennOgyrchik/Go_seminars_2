package storage

import (
	"testing"
)

func TestNewStorage(t *testing.T) {
	filenames := []string{"test1.txt", "test2", ".txt"}

	for i, filename := range filenames {
		_, err := NewStorage(filename)
		if err != nil {
			t.Fail()
			t.Logf(filenames[i], err)
		}
	}

}

func TestStorage_CreateUser(t *testing.T) {
	s, err := NewStorage("test.txt")
	if err != nil {
		t.Fail()
		t.Log()
	}

	user1 := User{
		Name:    "name1",
		Age:     50,
		Friends: nil,
	}
	user2 := User{
		Name:    "name2",
		Age:     50,
		Friends: []int{1},
	}
	Users := []User{user1, user2}

	for i, user := range Users {
		id := s.CreateUser(user)
		if id != i+1 {
			t.Fail()
			t.Log()
		}
	}

	if len(s.users) != 2 {
		t.Fail()
		t.Log()
	}
}

func TestStorage_MakeFriend(t *testing.T) {
	s, err := NewStorage("test.txt")
	if err != nil {
		t.Fail()
		t.Log()
	}

	id1 := []int{1, 2}
	id2 := []int{1, 1}

	user1 := User{
		Name:    "name1",
		Age:     50,
		Friends: nil,
	}
	user2 := User{
		Name:    "name2",
		Age:     50,
		Friends: nil,
	}

	s.users[1] = &user1
	s.users[2] = &user2

	for i := 0; i < 2; i++ {
		s.MakeFriend(id1[i], id2[i])
	}

	if len(s.users) != 2 {
		t.Fail()
		t.Log()
	}

}

func TestStorage_GetName(t *testing.T) {
	s, err := NewStorage("test.txt")
	if err != nil {
		t.Fail()
		t.Log()
	}

	user1 := User{
		Name:    "name1",
		Age:     50,
		Friends: nil,
	}
	user2 := User{
		Name:    "name2",
		Age:     50,
		Friends: nil,
	}

	s.users[1] = &user1
	s.users[2] = &user2

	id := []int{1, 2, 3}

	for _, i2 := range id {
		res, ok := s.GetName(i2)

		if res != "" && !ok {
			t.Fail()
			t.Log()
		}
	}
}

func TestStorage_DeleteUser(t *testing.T) {
	s, err := NewStorage("test.txt")
	if err != nil {
		t.Fail()
		t.Log()
	}

	user1 := User{
		Name:    "name1",
		Age:     50,
		Friends: []int{3},
	}
	user2 := User{
		Name:    "name2",
		Age:     50,
		Friends: []int{3},
	}
	user3 := User{
		Name:    "name3",
		Age:     50,
		Friends: []int{1, 2},
	}
	user4 := User{
		Name:    "name4",
		Age:     50,
		Friends: nil,
	}

	s.users[1] = &user1
	s.users[2] = &user2
	s.users[3] = &user3
	s.users[4] = &user4

	for i := 0; i < 4; i++ {
		s.DeleteUser(i)
		if len(s.users) != 4-i {
			t.Fail()
			t.Log()
		}
	}
}

func TestStorage_UpdateAge(t *testing.T) {
	s, err := NewStorage("test.txt")
	if err != nil {
		t.Fail()
		t.Log()
	}

	user1 := User{
		Name:    "name1",
		Age:     50,
		Friends: []int{3},
	}
	user2 := User{
		Name:    "name2",
		Age:     50,
		Friends: []int{3},
	}
	user3 := User{
		Name:    "name3",
		Age:     50,
		Friends: []int{1, 2},
	}
	user4 := User{
		Name:    "name4",
		Age:     50,
		Friends: nil,
	}

	s.CreateUser(user1)
	s.CreateUser(user2)
	s.CreateUser(user3)
	s.CreateUser(user4)
	ages := []int{1, 55, 87, 2126}

	for i := 0; i < 4; i++ {
		s.UpdateAge(i+1, ages[i])
		if s.users[i+1].Age != ages[i] {
			t.Fail()
			t.Log()
		}
	}
}
