package _interface

import "testing"

// 定义一个接口 (只要方法签名是一样的就是实现了接口)
type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (t *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"hello,world\")"
}

func TestClient(t *testing.T) {
	var p Programmer

	p = new(GoProgrammer)

	t.Log(p.WriteHelloWorld())

}
