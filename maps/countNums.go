package main

import (
	"fmt"
)

func main() {
	list := []int{1, 2, 3, 4, 1, 2, 4}
	result := countNums(list)
	fmt.Println("result", result)
}

// 统计切片中每个整数出现的次数
func countNums(arr []int) map[int]int {
	result := map[int]int{}
	for _, i2 := range arr {
		result[i2]++
	}
	return result
}
