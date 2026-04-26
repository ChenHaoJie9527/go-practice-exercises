package main

import "testing"

func TestReverseArray(t *testing.T) {
	arr := [3]int{1, 2, 3}
	reverseArray(&arr)
	if arr != [3]int{3, 2, 1} {
		t.Errorf("反转后的数组应该是[3, 2, 1], 实际是: %v", arr)
	}
}