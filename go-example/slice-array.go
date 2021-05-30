package main

import "fmt"

func main() {
	var arr = []int{1, 2, 3, 4, 5, 6}
	var lenArr = len(arr)
	for i := 0; i < lenArr; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()

	var slice []int
	printSlice(slice)

	if (slice == nil) {
		fmt.Println("slice is nil")
	}

	var nums = make([]int, 10, 20)
	fmt.Println("nums: ");
	printSlice(nums)
	var nums1 []int
	nums1 = nums
	nums = append(nums, 10)
	nums1[1] = 111
	fmt.Printf("nums1:");
	printSlice(nums1)
	fmt.Println("nums:")
	printSlice(nums)

	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	printSlice(numbers)

	printSlice(numbers[1:6])

	printSlice(numbers[:5])

	printSlice(numbers[2:])


	numbers1 := append(numbers, 8)
	numbers[1] = 9

	printSlice(numbers)
	printSlice(numbers1)

	numbers1 = append(numbers1, 9 ,10, 11)
	printSlice(numbers1)

	finalNumbers := make([]int, len(numbers1), 100)
	copy(finalNumbers, numbers1)

	printSlice(finalNumbers)




}

func printSlice(s []int) {
   fmt.Printf("len = %d cap = %d slice = %v\n", len(s), cap(s), s)
}

