package main

type User struct {
	ID                         int
	FirstName, LastName, Email string
}

func updateEmail(user *User, newEmail string) User {
	user.Email = newEmail
	return *user
}

func main() {
	user := User{ID: 1, FirstName: "Necip", LastName: "AKGZ", Email: "fake@gmail.com"}

	updateEmail(&user, "nex@hotmail.com")
	println(user.Email)
}
