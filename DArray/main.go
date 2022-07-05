package main

import (
	"errors"
	"fmt"
	"strconv"
)

const(
	growCoefficient = 2
	startSize = 1
)

type DArray struct {
	buffer []int
	size int
	capacity int
}

func (darray *DArray) Add (newElement int) {

	if (darray.size == darray.capacity || darray.capacity == 0){
		darray.Grow()
	}
	darray.buffer[darray.size] = newElement
	darray.size++
}

func (darray *DArray) Grow () {
	if darray.capacity == 0{
		darray.buffer = make([]int, startSize)
		darray.capacity = startSize
		return
	}
	oldBuffer := darray.buffer
	darray.buffer = make([]int, darray.capacity * growCoefficient)
	copy(darray.buffer, oldBuffer)
	darray.capacity = darray.capacity * growCoefficient
}

func (darray *DArray) Size() (int){
	return darray.size
}

func (darray *DArray) GetAt(index int) (int, error){
	if(index > darray.size - 1){
		return 0, errors.New("Out of size!")
	}
	return darray.buffer[index], nil
}

func (darray *DArray) Print() {
	fmt.Println("DArray size = " + strconv.Itoa(darray.Size()))
	fmt.Println("DArray capacity = " + strconv.Itoa(darray.capacity))
	fmt.Println("DArray elements:")
	for i := 0; i < darray.size; i++ {
		fmt.Println(darray.buffer[i])
	}

}

func main () {

}