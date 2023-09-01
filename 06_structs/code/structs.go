package main

type User struct {
	ID                    int
	name, lastname, email string
}

func main() {
	var user User = User{ID: 1, name: "Necip", lastname: "AKGOZ"}
	println(user.name)
}
