package main

import "fmt"

type User struct {
	username string
	Password string
	Uname    string
}

func main() {
	var user User
	user.username = "zhangsan"
	fmt.Printf("username:%v \n", user.username)
	fmt.Println("hello go")
}
