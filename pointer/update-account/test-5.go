package main

import "fmt"

type Account struct {
	Owner string
	Balance float64 // 使用float64 而不是int，因为余额可能有小数，不实用float32的原因是float32的精度不够，容易出现精度丢失问题。
}

// 声明deposit方法，用于存款
// 为什么要 (a *Account) 呢？为什么要这样写，有什么意义？
// 解答：
// (a *Account）表示方法指针接受者，传递的是内存地址，方法里改动 直接修改原对象
// (a Account) 表示方法值接受者，传递的是值的拷贝，方法里改动不会影响原对象
func (a *Account) Deposit(amount float64) {
	if a == nil {
		fmt.Println("账户不存在")
		return
	}

	a.Balance += amount
}