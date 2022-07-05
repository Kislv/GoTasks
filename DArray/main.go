package DArray

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
		darray.grow()
	}
	darray.buffer[darray.size] = newElement
	darray.size++
}

func (darray *DArray) grow () {
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

func (darray *DArray) Delete () (error) {

	if darray.Size() == 0{
		return errors.New("DArray is empty! ")
	}
	darray.size--
	if (darray.size == darray.capacity / 2){
		err := darray.shrink()
		if err != nil {
			return err
		}
	}

	return nil
}

func (darray *DArray) shrink () (error){
	if darray.capacity == 0 {
		return errors.New("darray.capacity == 0")
	}
	oldBuffer := darray.buffer
	darray.buffer = make([]int, darray.capacity / growCoefficient)
	copy(darray.buffer, oldBuffer)
	darray.capacity = darray.capacity / growCoefficient
	return nil
}

func (darray *DArray) Size() (int){
	return darray.size
}

func (darray *DArray) GetAt(index int) (*int, error){
	if(index > darray.size - 1){
		return nil, errors.New("Out of size! ")
	}
	return &darray.buffer[index], nil
}
func (darray *DArray) IsEmpty() (bool){
	return darray.Size() == 0 
}

func (darray *DArray) Print() {
	fmt.Println("DArray size = " + strconv.Itoa(darray.Size()))
	fmt.Println("DArray capacity = " + strconv.Itoa(darray.capacity))
	fmt.Println("DArray elements:")
	for i := 0; i < darray.size; i++ {
		fmt.Println(darray.buffer[i])
	}

}
