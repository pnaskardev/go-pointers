package main

import "fmt"

// when to use
// - when we need to update the state
// - when we want to optimise the memory for large objects that re getting called a lot


type User struct {
	email    string
	username string
	age      int
}

func (u User) Email() string {
	return u.email
}

func (u *User) updateEmail(email string) {
	u.email = email

}
func main() {

	user := User{
		email: "pnaskrdev@gmail.com",
	}
	fmt.Println(user.Email())

	user.updateEmail("test@test.com")

	fmt.Println(user.Email())

}
