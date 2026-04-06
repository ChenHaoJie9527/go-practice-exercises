package main

import "fmt"

func main() {
	fmt.Println(isAnagram("abc", "bca")) // true
	fmt.Println(isAnagram("aab", "aba")) // true
	fmt.Println(isAnagram("aab", "abb")) // false
}

func isAnagram(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	count := make(map[byte]int)

	for i := 0; i < len(str1); i++ {
		count[str1[i]]++
	}

	for i := 0; i < len(str2); i++ {
		count[str2[i]]--
	}

	for _, v := range count {
		if v != 0 {
			return false
		}
	}
	return true

}
