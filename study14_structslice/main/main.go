package main

import "fmt"

type User struct {
	ID    uint
	Name  string
	Uname string
}

type UserDto struct {
	ID   uint
	Name string
}

func GetUsers(user User) []UserDto {
	
	return []UserDto{
		{
			user.ID,
			user.Name,
		},
	}
}

func main() {
	var user User
	userdto := GetUsers(user)

	fmt.Printf("%v",userdto)
}