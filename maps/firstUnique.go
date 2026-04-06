package main

import "fmt"

func main() {
	nums := []int{4, 5, 1, 2, 1, 2, 5}

	result, ok := firstUnique(nums)
	if ok {
		fmt.Println("第一个只出现一次的数字是:", result)
	} else {
		fmt.Println("没有只出现一次的数字")
	}
}

func firstUnique(arr []int) (int, bool) {
	counts := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		counts[arr[i]]++
	}

	for i := 0; i < len(arr); i++ {
		if counts[i] == 1 {
			return i, true
		}
	}

	return 0, false
}
