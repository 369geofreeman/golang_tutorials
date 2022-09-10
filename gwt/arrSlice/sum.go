package main

func Sum(numbers []int) int {
	var total int
	for _, x := range numbers {
		total += x
	}
	return total
}
