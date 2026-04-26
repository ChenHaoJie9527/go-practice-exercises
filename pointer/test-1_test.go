package main

import (
	"testing"
)

func TestZeroval_DoesNotMutateCaller(t * testing.T) {
	i := 1
	zeroval(i)

	if i != 1 {
		t.Errorf("zeroval 不应该修改调用者的值: %d", i)
	}
}

func TestZeroptr_MutatesCaller(t * testing.T) {
	i := 1
	zeroptr((&i))

	if i != 0 {
		t.Errorf("zeroptr 应该修改调用者的值: %d", i)
	}
}

func TestZeroptr_ReturnsPointer(t * testing.T) {
	tests := []struct {
		name string
		input int
	} {
		{name: "正数", input: 42},
		{name: "负数", input: -7},
		{name: "零值", input: 0},
	}

	for _, v := range tests {
		t.Run(v.name, func(t * testing.T) {
			i := v.input
			zeroptr(&i)
			if i != 0 {
				t.Errorf("zeroptr 应该返回指针: %d", i)
			}
		})
	}

}