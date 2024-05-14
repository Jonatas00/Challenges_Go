package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 26))
}

func twoSum(nums []int, target int) []int {
	numerosPercorridos := make(map[int]int)
	for k, v := range nums {
		complemento := target - v

		_, existsInNumerosPercorridos := numerosPercorridos[complemento]
		if existsInNumerosPercorridos {
			return []int{k, numerosPercorridos[complemento]}
		}
		numerosPercorridos[v] = k
	}
	return []int{}
}
