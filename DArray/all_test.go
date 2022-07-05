package DArray

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFill(t *testing.T) {
	var darray DArray
	for i := 0; i < 5; i++ {
		darray.Add(i)
	}
	element, err := darray.GetAt(4)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(*element, 4) {
		t.Error("DArray are not equal")
	}
}

func TestDelete(t *testing.T) {
	var darray DArray
	for i := 0; i < 5; i++ {
		darray.Add(i)
	}
	for i := 0; i < 3; i++{
		err := darray.Delete()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	darray.Add(100)
}