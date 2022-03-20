package usercollection

import (
	"github.com/michelo851a1203/testiter/iterator"
	"github.com/michelo851a1203/testiter/useriterator"
	"github.com/michelo851a1203/testiter/usertype"
)

type UserCollection struct {
	UserGroup []*usertype.UserType
}

func (userCollection *UserCollection) CreateIterator() iterator.Iterator {
	return &useriterator.UserIterator{
		UserGroup: userCollection.UserGroup,
	}
}
