package main

import "testing"

func TestModifySlice(t *testing.T) {
	slice := []int{1, 2, 3}
	modifySlice(slice)
	if slice[0] != 100 {
		t.Errorf("修改后的切片应该是[100, 2, 3], 实际是: %v", slice)
	}
}