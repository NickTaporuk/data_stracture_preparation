package LinkedList

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleLinkedList_Insert(t *testing.T) {
	l := SingleLinkedList{}

	err := l.Insert(0, 1)
	assert.NoError(t, err)

	err = l.Insert(1, 2)
	assert.NoError(t, err)

	err = l.Insert(2, 3)
	assert.NoError(t, err)

	l.Append(4)
	l.Append(5)
	l.Append(6)
	//
	err = l.Insert(2, 7)
	assert.NoError(t, err)

	_, err = l.RemoveAt(1)
	assert.NoError(t, err)
	l.Remove()
	l.Remove()
	looked := l.IndexOf(5)
	assert.Equal(t, 4, looked)
	fmt.Println(l)
}

func TestSingleLinkedList_Reverse(t *testing.T) {
	l := Constractor()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Append(7)
	l.Append(8)

	l.Reverse()

	fmt.Println(l)
}

func TestMergeSorted(t *testing.T) {
	l := Constractor()

	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)

	l2 := Constractor()

	l2.Append(1)
	l2.Append(2)
	l2.Append(3)
	l2.Append(4)

	l3 := Constractor()
	var err error

	l3.Head, err = MergeSorted(l.Head, l2.Head, isSortedAsc)
	assert.NoError(t, err)


	fmt.Println(l3)

}