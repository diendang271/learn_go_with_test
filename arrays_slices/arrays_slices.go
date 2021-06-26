package arrays_slices

func Sum(numbers [5]int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumV2(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}


