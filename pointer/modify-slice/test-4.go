package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	modifySlice(slice) // 切片不需要传递指针，因为切片本身就是指针
	fmt.Println(slice)
}

func modifySlice(slice []int) {
	slice[0] = 100
}