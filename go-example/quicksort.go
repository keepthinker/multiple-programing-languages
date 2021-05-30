package main

import "fmt"
import "time"

func main() {
	var arr = []int {10, 2, 4, 8, 20, 100, 20, 9, 3, 1, 38, 2, 1}
	length := len(arr)
	fmt.Println("initial arr:")
	printArr(arr, 0, length - 1, length - 1)

	quicksort(arr, 0, length - 1);
	for _, value := range arr  {
		fmt.Printf("%d ", value)
	}
	fmt.Println();
}

func printArr(arr []int, begin int, end int, pivotIndex int) {
	fmt.Printf("---------\npivotIndex %d, begin %d, end %d\n", pivotIndex, begin, end)
	for i,value := range arr{
		fmt.Printf("%d:%d  ", i, value)
	}
	fmt.Println("\n---------");
}

func quicksort(arr []int, begin int, end int) {
	if (begin >= end) {
		return
	}
	pivotIndex := partion(arr, begin, end)

	printArr(arr, begin, end, pivotIndex)
	time.Sleep(100 * time.Millisecond)
	quicksort(arr, begin, pivotIndex - 1)
	quicksort(arr, pivotIndex + 1, end)
}

func partion(arr []int, begin int, end int) int {
	pivotVal := arr[end]
	index := begin - 1;
	i := begin
	for i < end {
		if arr[i] < pivotVal {
			temp := arr[i]
			index++
			arr[i] = arr[index]
			arr[index] = temp
		}
		i++;
	}
	index++
	temp := arr[index]
	arr[index] = pivotVal
	arr[end] = temp
	return index
}
