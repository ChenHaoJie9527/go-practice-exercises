package main

import (
	"fmt"
)

func main() {
	str := "aabbcedsfsen"
	counts := countCharacters(str)
	for k, v := range counts {
		fmt.Printf("%c: %d\n", k, v)
	}

	str1 := "你好啊好好你你说奋斗的少年"
	counts1 := countCharacters2(str1)
	for k, v := range counts1 {
		fmt.Printf("%c: %d\n", k, v)
	}
}

// 统计字符串中每个字符出现的次数
// 要求：输入一个字符串，返回每个字符出现次数的 map。先只考虑英文和数字。
// 提示：遍历字符串时，先想清楚你拿到的是字节还是字符。
// 思考方向：map 的 key 应该用什么类型；不存在的 key 读取出来是什么值；如何做计数累加。
func countCharacters(str string) map[byte]int {
	result := make(map[byte]int)

	for i := 0; i < len(str); i++ {
		result[str[i]]++
	}

	return result
}

func countCharacters2(str string) map[rune]int {
	result := make(map[rune]int)
	for _, v := range str {
		result[v]++
	}

	return result
}
