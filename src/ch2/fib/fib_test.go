package fib

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T) {
	//var a int = 1
	//var b int = 2

	//var (
	//	a int = 1
	//	b int = 2
	//)

	a := 1
	b := 1
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}

	fmt.Println()
}

func TestExchange(t *testing.T)  {
	a := 1
	b := 6

	a,b=b, a

	t.Log(a, b)
}
