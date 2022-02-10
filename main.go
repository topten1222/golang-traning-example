package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5}
	result, err := tailSum(array, 20)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func tailSum(value []int, total int) (int, error) {
	length := len(value)
	result := 0
	countLoop := 0
	for i := length; i > 0; i-- {
		countLoop++
		key := i - 1
		result += value[key]
		if result == total {
			return countLoop, nil
		}
	}
	return 0, nil
}
