package main

import "fmt"

func main() {

	nums := []int{1, 10, 30, 40, 50, 60, 70}
	item := 30

	chute, qtd_operacoes := BinSearch(nums, item)
	fmt.Printf("chute: %v qtd_op: %v\n", chute, qtd_operacoes)

}

func BinSearch(nums []int, item int) (int, int) {

	baixo := 0
	alto := len(nums) - 1
	count := 0

	for baixo <= alto {
		meio := (alto + baixo) / 2
		chute := nums[meio]
		count++

		fmt.Printf("baixo: %v - %v ## meio: %v - %v ## alto: %v - %v \n", baixo, nums[baixo], meio, nums[meio], alto, nums[alto])

		if chute == item {
			return chute, count
		}
		if chute > item {
			alto = meio - 1
		} else {
			baixo = meio + 1
		}
	}
	return 0, 0
}
