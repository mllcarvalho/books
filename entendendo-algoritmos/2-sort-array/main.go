package main

import "fmt"

func main() {

	nums := []int{5, 10, 3, 2, 1}
	fmt.Println(Sort(nums))

}

func Sort(nums []int) []int {

	for i := len(nums); i > 0; i-- {
		for j := 1; j < i; j++ {
			if nums[j-1] > nums[j] {
				aux := nums[j]
				nums[j] = nums[j-1]
				nums[j-1] = aux
			}
		}
	}
	return nums
}

// package main

// import "fmt"

// func main() {

// 	nums := []int{5, 4, 3, 2, 1}
// 	fmt.Println(sortSlice(nums))
// }

// func sortSlice(S []int) []int {
// 	for i := len(S); i > 0; i-- {
// 		for j := 1; j < i; j++ {
// 			fmt.Println(S[j-1], S[j])
// 			if S[j-1] > S[j] {
// 				// swap
// 				intermediate := S[j]
// 				S[j] = S[j-1]
// 				S[j-1] = intermediate
// 			}
// 			fmt.Println("swap", S[j-1], S[j])
// 			fmt.Println("array", S)
// 		}
// 		fmt.Println("array for", S)
// 	}
// 	return S
// }
