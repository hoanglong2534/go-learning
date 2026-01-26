package _struct

import "fmt"

type User struct {
	firstName string
	lastName  string
	age       int
}

func MyStruct() {

	listUser := []User{}

	user1 := User{
		firstName: "Hoang",
		lastName:  "Long",
		age:       21,
	}

	admin := User{
		firstName: "admin",
		lastName:  "Pham",
		age:       22,
	}

	listUser = append(listUser, user1, admin)

	for i := 0; i < len(listUser); i++ {
		firstName := listUser[i].firstName
		lastName := listUser[i].lastName
		age := listUser[i].age
		fmt.Println(i, firstName, lastName, age)
	}
}
