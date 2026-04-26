package main

import "fmt"

// reverseArray 原地反转一个长度为 3 的整数数组。
// 参数使用 *[3]int 是为了避免数组按值拷贝，并让函数能直接修改调用方的数组。
func reverseArray(arr *[3]int) {
	left := 0
	right := len(arr) - 1

	// 循环 条件是 left < right
	for left < right {
		temp := arr[left] // 取最左边元素的值
		arr[left] = arr[right] // 将最右边元素的值赋给最左边元素
		arr[right] = temp // 将最左边元素的值赋给最右边元素
		left++ // 递增+1
		right-- // 递减-1
	}
}

func main() {
	arr := [3]int{1, 2, 3}
	reverseArray(&arr)
	fmt.Println("反转后的数组", arr)
	
}