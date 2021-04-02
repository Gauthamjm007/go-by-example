package array

import (
	"errors"
	"fmt"
)

// Creating a Structhre for Array of type Int
type Array struct {
	data   []int
	length int
}

// function to Initialize an Array
func NewArray(capacity uint) *Array {
	if capacity == 0 {
		return nil
	}
	return &Array{
		data:   make([]int, capacity),
		length: 0,
	}
}

// Method on Array to find it's length
func (a *Array) Len() int {
	return a.length
}

// Method on Array to find if the index is out of range
func (a *Array) isIndexOutofRange(index uint) bool {
	intIndex := int(index)
	return intIndex >= a.length
}

// Methods on array to Find an element in Array by index and return the value

func (a *Array) Find(index uint) (int, error) {
	if a.isIndexOutofRange(index) {
		return 0, errors.New("out of index range")
	}
	return a.data[index], nil
}

// Method on Array to Insert an Element to a specific Index

func (a *Array) Insert(index uint, v int) error {
	if uint(a.Len()) == uint(cap(a.data)) {
		return errors.New("full array")
	}

	if index != uint(a.length) && a.isIndexOutofRange(index) {
		return errors.New("out of  index range")
	}

	for i := a.length; uint(i) > index; i-- {
		a.data[i] = a.data[i-1]
	}
	a.data[index] = v
	a.length++
	return nil
}

// Methods on Array to Push an int to end of the Array

func (a *Array) InsertTailend(v int) error {
	return a.Insert(uint(a.Len()), v)
}

//Methods on Array to Print the Array

func (a *Array) Print() {
	var format string
	for i := 0; i < a.Len(); i++ {
		format += fmt.Sprintf("|%+v", a.data[i])
	}
	fmt.Println(format)
}
