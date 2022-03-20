package useriterator

import "github.com/michelo851a1203/testiter/usertype"

type UserIterator struct {
	Index     int
	UserGroup []*usertype.UserType
}

func (userIterator *UserIterator) GetNext() *usertype.UserType {
	if userIterator.HasNext() {
		currentUserNode := userIterator.UserGroup[userIterator.Index]
		userIterator.Index++
		return currentUserNode
	}
	return nil
}

func (userIterator *UserIterator) HasNext() bool {
	userGroupLen := len(userIterator.UserGroup)
	if userIterator.Index < userGroupLen {
		return true
	}
	return false
}
