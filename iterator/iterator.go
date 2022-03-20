package iterator

import "github.com/michelo851a1203/testiter/usertype"

type Iterator interface {
	GetNext() *usertype.UserType
	HasNext() bool
}
