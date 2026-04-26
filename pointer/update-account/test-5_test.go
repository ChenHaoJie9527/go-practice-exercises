package main

import (
	"fmt"
	"testing"
)

func TestDeposit(t *testing.T) {
	accounts := []*Account{
		{
			"Alice",
			100,
		},
		{
			"Bob",
			200,
		},
	}

	for _, v := range accounts {

		if v.Owner == "Alice" {
			v.Deposit(200)
		}
	}

	for _, v := range accounts {
		fmt.Printf("%s balance: %.2f\n", v.Owner, v.Balance)
		if v.Owner == "Alice" && v.Balance != 300 {
			t.Errorf("Alice 的余额应该是300, 实际是: %.2f", v.Balance)
		}
	}

}
