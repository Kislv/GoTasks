package Heap

import (
	"GoSTL/DArray"
	"fmt"
	"reflect"
	"testing"
)

func TestBuild(t *testing.T) {
	var heap Heap
	heap.array.Add(3)
	heap.array.Add(7)
	heap.array.Add(9)
	heap.array.Add(6)
	heap.array.Add(8)
	heap.array.Add(4)
	err := heap.BuildHeap()
	if err != nil {
		fmt.Println(err)
		return
	}

	expectedDArray := DArray.DArray{}
	expectedDArray.Add(9)
	expectedDArray.Add(8)
	expectedDArray.Add(4)
	expectedDArray.Add(6)
	expectedDArray.Add(7)
	expectedDArray.Add(3)
	if !reflect.DeepEqual(heap.array, expectedDArray) {
		t.Error("Heap build error")
	}

}

func TestPeek(t *testing.T) {
	var heap Heap
	heap.Insert(3)
	heap.Insert(7)

	max, err := heap.PeekMax()
	if err != nil {
		fmt.Println(err)
		return
	}
	if !reflect.DeepEqual(max, 7) {
		t.Error("TestPeek error")
	}

}


func TestExtract(t *testing.T) {
	var heap Heap
	heap.Insert(3)
	heap.Insert(7)
	heap.Insert(9)
	heap.Insert(6)
	heap.Insert(8)
	heap.Insert(4)

	max, err := heap.ExtractMax()
	if err != nil {
		fmt.Println(err)
		return
	}
	if !reflect.DeepEqual(max, 9) {
		t.Error("TestExtract error")
	}

	max, err = heap.ExtractMax()
	if err != nil {
		fmt.Println(err)
		return
	}
	if !reflect.DeepEqual(max, 8) {
		t.Error("TestExtract error")
	}

}
