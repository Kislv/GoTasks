package Heap

import (
	"GoSTL/DArray"
	"errors"
)

type Heap struct{
	array DArray.DArray
}

func (heap *Heap) Insert (in int){
	heap.array.Add(in)
	heap.siftUp(heap.array.Size() - 1)
}


func (heap *Heap) ExtractMax () (int, error){
	if heap.array.IsEmpty() {
		return 0, errors.New("Out of size! ")
	}
	pMax, err := heap.array.GetAt(0)
	if err != nil {
		return 0, err
	}
	max := *pMax
	pLast, err := heap.array.GetAt(heap.array.Size() - 1)
	if err != nil {
		return 0, nil
	}
	
	*pMax = *pLast
	err = heap.array.Delete()
	if err != nil {
		return 0, err
	}
	err = heap.siftDown(0)
	if err != nil {
		return 0, err
	}

	return max, nil
}


func (heap *Heap) PeekMax () (int, error){
	if heap.array.IsEmpty() {
		return 0, errors.New("Out of size! ")
	}
	pMax, err := heap.array.GetAt(0)	
	if err != nil {
		return 0, err
	}
	return *pMax, nil
}

func (heap *Heap) BuildHeap () (error){
	for i := heap.array.Size() / 2 - 1; i >= 0; i-- {
		err := heap.siftDown(i)
		if err != nil {
			return err
		}
	} 

	return nil
}

func (heap *Heap) siftDown (index int) (error){
	if index > heap.array.Size() - 1 {
		return errors.New("SiftDown: out of size")
	}
	size := heap.array.Size()
	leftIndex := 2 * index + 1
	rightIndex := 2 * index + 2
	
	currentPointer, err := heap.array.GetAt(index)
	if(err != nil){
		return err
	}
	currentValue := *currentPointer

	maxChildIndex := index
	maxChildValue := currentValue

	if (leftIndex < size) {
		pLeft, err := heap.array.GetAt(leftIndex)
		if (err != nil){
			return err
		}
		if *pLeft > currentValue{
			maxChildIndex = leftIndex
			maxChildValue = *pLeft

		}
	} else {
		return nil
	}

	if(rightIndex < size) {
		pRight, err := heap.array.GetAt(rightIndex)
		if (err != nil){
			return err
		}
		if *pRight > maxChildValue{
			maxChildIndex = rightIndex
		}
	}

	

	maxChildP, err := heap.array.GetAt(maxChildIndex)
	if(err != nil){
		return err
	}

	if (currentValue >= *maxChildP){
		return  nil
	} else {
		*currentPointer, *maxChildP = *maxChildP, *currentPointer
		err = heap.siftDown(maxChildIndex)
		if err != nil {
			return err
		} 
	}
		
	return nil
}


func (heap *Heap) siftUp (index int) (error){
	if index == 0 {
		return nil
	}

	if index > heap.array.Size() - 1 {
		return errors.New("siftUp: out of size")
	}

	currentP, err := heap.array.GetAt(index)
	if err != nil {
		return err
	}
	current := *currentP

	parentIndex := (index - 1) / 2
	parentP, err := heap.array.GetAt(parentIndex)
	if err != nil {
		return err
	}
	parent := *parentP

	if current > parent {
		*currentP, *parentP = *parentP, *currentP
		heap.siftUp(parentIndex)
	}
	return nil
}
	
func maxInt (first, second int) (bool){
	return first >= second
} 

func main () {

}