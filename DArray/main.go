package main

import(

)

type DArray struct {
	buffer []int
	size int
	capacity int
}

func (darray *DArray) Add (int) {
	if (darray.size == darray.capacity){
		darray.Grow()
	}
}

func (darray *DArray) Grow () {
	oldBuffer := darray.buffer
	darray.buffer = 
}

// READ ABOUT ARRAYS!

func main () {

}