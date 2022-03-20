package main

import (
	"fmt"

	"github.com/michelo851a1203/testiter/usercollection"
	"github.com/michelo851a1203/testiter/usertype"
)

func main() {
	userOne := usertype.UserType{
		Name:  "test01",
		Email: "test01@gmail.com",
		Age:   12,
	}

	userTwo := usertype.UserType{
		Name:  "test02",
		Email: "test02@gmail.com",
		Age:   23,
	}

	myCollection := usercollection.UserCollection{
		UserGroup: []*usertype.UserType{&userOne, &userTwo},
	}
	myIterator := myCollection.CreateIterator()

	for myIterator.HasNext() {
		userNode := myIterator.GetNext()
		fmt.Println(userNode)
	}
}
