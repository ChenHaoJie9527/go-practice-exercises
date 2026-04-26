package main

import (
	"fmt"
	"testing"
)

func TestUpdateAge(t *testing.T) {
	p1 := Person{
		"张三",
		18,
	}

	updateAge(&p1, 20)

	if p1.age != 20 {
		t.Errorf("更新后的年龄应该是20, 实际是: %d", p1.age)
	}

	p2 := new(Person)
	fmt.Println("p2初始零值", p2)
	p2.name = "Bob"
	p2.age = 25
	fmt.Println("p2更新后的值", p2)

	updateAge(p2, 30)

	if p2.age != 30 {
		t.Errorf("更新后的年龄应该是30, 实际是: %d", p2.age)
	}
}