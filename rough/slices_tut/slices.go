package main

// slices -> dynamic
// most used in go
func main() {
	// uninitialised slice is nil
	// var nums []int

	// fmt.Println(nums)
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	// var nums = make([]int,2,5)
	// max number of elements can fit
	// fmt.Println(cap(nums))

	// nums := []int{}

	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 2)
	// nums = append(nums, 2)
	// nums = append(nums, 2)
	// nums = append(nums, 2)
	// nums = append(nums, 2)
	// nums = append(nums, 2)

	// fmt.Println(nums)
	// fmt.Println(cap(nums))

	// var nums = make([]int,0,5)
	// nums = append(nums, 1)

	// var nums_copy = make([]int,len(nums))

	// copy(nums_copy,nums)

	// fmt.Println(nums,nums_copy)

	// Slice operator

	// var nums = []int{1,2,3}

	// fmt.Println(nums[0:2])

	// Slice package

	// var nums_1 = []int{1,2}
	// var nums_2 = []int{1,2,4}

	// fmt.Println(slices.Equal(nums_1,nums_2))

	// 2D slices

	// var nums = [][]int{{1, 2, 3}, {2, 3, 4}}
	// fmt.Println(nums)

}
