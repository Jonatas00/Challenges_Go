package main

// Return a new array consisting of elements which are multiple of their own
// index in input array (length > 1).

func main() {
	multipleOfIndex([]int{22, -6, 32, 82, 9, 25})
	multipleOfIndex([]int{68, -1, 1, -7, 10, 10})
}

func multipleOfIndex(ints []int) (multiples []int) {
	for i, v := range ints {
		if i > 0 && v%i == 0 {
			multiples = append(multiples, v)
		}
	}
	return multiples
}
