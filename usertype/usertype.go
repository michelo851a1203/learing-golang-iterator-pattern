package usertype

import "fmt"

type UserType struct {
	Name  string
	Email string
	Age   int
}

func (userType UserType) String() string {
	return fmt.Sprintf("=====\n\n Name : %s\n Email : %s\n Age : %d\n\n=====",
		userType.Name,
		userType.Email,
		userType.Age,
	)

}
