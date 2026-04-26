package main

import "fmt"

type Person struct {
	name string
	age int
}

func updateAge(per *Person, newAge int) {
	per.age = newAge
}

func main() {
	// 使用字面量创建结构体实例
	p1 := Person{
		"张三",
		18,
	}

	// 使用 new 关键字创建结构体实例，但不方便初始化实例
	p2 := new(Person)
	fmt.Println(p2)
	p2.name = "莉丝"
	p2.age = 25
	fmt.Println(p2)

	updateAge(&p1, 20)
	// 使用 new 关键字创建结构体实例，可以直接传递指针
	updateAge(p2, 30)


	// --------------------
	// new(T): 在内存中申请一块能放下 T 类型实例的空间，并返回这个空间的地址，即 *T 类型的值。
	// 把这块空间默认设置为 T 类型实例的零值。
}

