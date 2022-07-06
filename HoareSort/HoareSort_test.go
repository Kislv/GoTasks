package HoareSort

import (
	"reflect"
	"testing"
)

func TestFewElements(t *testing.T) {
	slice := []int{5,1,3}
	HoareSort(slice)
	expectedSlice := []int{1,3,5}
	if !reflect.DeepEqual(slice, expectedSlice) {
		t.Error("TestFewElements failed")
	}
}

func TestNoSort(t *testing.T) {
	slice := []int{-5, 0, 0, 5, 10, 10, 10}
	HoareSort(slice)
	expectedSlice := []int{-5, 0, 0, 5, 10, 10, 10}
	if !reflect.DeepEqual(slice, expectedSlice) {
		t.Error("TestNoSort failed")
	}
}


func TestWorstCase(t *testing.T) {
	slice := []int{1000, 999, 998, 998, 10, 10, 10, 5, 2}
	HoareSort(slice)
	expectedSlice := []int{2, 5, 10, 10, 10, 998, 998, 999, 1000}
	if !reflect.DeepEqual(slice, expectedSlice) {
		t.Error("TestWorstCase failed")
	}
}


func Test8Elements(t *testing.T) {
	slice := []int{10, -10, 4, 4, -2, 8, 0 , 6}
	HoareSort(slice)
	expectedSlice := []int{-10, -2, 0, 4, 4, 6, 8, 10}
	if !reflect.DeepEqual(slice, expectedSlice) {
		t.Error("Test8Elements failed")
	}
}

func TestEmpty(t *testing.T) {
	slice := []int{}
	HoareSort(slice)
	expectedSlice := []int{}
	if !reflect.DeepEqual(slice, expectedSlice) {
		t.Error("TestEmpty failed")
	}
}

func Test1Element(t *testing.T) {
	slice := []int{6}
	HoareSort(slice)
	expectedSlice := []int{6}
	if !reflect.DeepEqual(slice, expectedSlice) {
		t.Error("Test1Element failed")
	}
}

func Test2Element(t *testing.T) {
	slice := []int{6, 3}
	HoareSort(slice)
	expectedSlice := []int{3, 6}
	if !reflect.DeepEqual(slice, expectedSlice) {
		t.Error("Test2Element failed")
	}
}

func TestSameElements(t *testing.T) {
	slice := []int{3, 3, 3, 3, 3}
	HoareSort(slice)
	expectedSlice := []int{3, 3, 3, 3, 3}
	if !reflect.DeepEqual(slice, expectedSlice) {
		t.Error("TestSameElements failed")
	}
}
