package collection

import "github.com/michelo851a1203/testiter/iterator"

type Collection interface {
	CreateIterator() iterator.Iterator
}
