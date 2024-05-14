package main

func main() {
	twoSum([]int{2, 7, 11, 15}, 26)
}

func twoSum(nums []int, target int) []int {
	//numerosPercorridos := make(map[int]int)

	for k1, v1 := range nums {
		for k2, v2 := range nums {
			if v1+v2 == target {
				return []int{k1, k2}
			}
		}
	}
	return []int{}
}
