package main

import (
	"fmt"
	"strconv"
)

func main() {
	result, err := parseAndAdd("10", "20")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(result)
}

func parseAndAdd(a, b string) (int, error) {
	x, err := strconv.Atoi(a)
	fmt.Println("x: ", x, "err:", err)
	// 判断 err 是否有值
	if err != nil {
		return 0, err
	}
	// 判断 err 是否有值
	y, err := strconv.Atoi(b)
	if err != nil {
		return 0, err
	}

	return x + y, nil
}
