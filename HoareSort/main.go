package HoareSort

import(

)

func partition (slice []int) (int) {
	sliceSize := len(slice)
	pivotIndex := sliceSize - 1
	pivot := slice[pivotIndex]
	i := 0  
	k := sliceSize - 2 
	for i <= k {
		if i == k {
			if (slice[k] > slice[pivotIndex]) {
				slice[pivotIndex], slice[k] = slice[k], slice[pivotIndex]
			}
			return k
		}
		for i < pivotIndex {
			if (slice[i] >= pivot) {
				break
			}
			i++
		}
		for k > 0 {
			if slice[k] < pivot && i < k {
				slice[i], slice[k] = slice[k], slice[i]
				break
			}
			k--
		}
	}
	slice[i], slice[pivotIndex] = slice[pivotIndex], slice[i]
	return i
}

func HoareSort (slice []int) {
	sliceSize := len(slice)
	if (sliceSize  <= 1){
		return
	}
	edge := partition(slice)
	if (edge != 0 ){
		HoareSort(slice[:edge])
	}
	if (edge != sliceSize - 1) {
		HoareSort(slice[edge + 1:])
	}
}