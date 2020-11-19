package empty_interface

import (
	"fmt"
	"testing"
)

func doSomething(p interface{}) {
	//if i,ok:=p.(int);ok{
	//	fmt.Println("integer", i)
	//	return
	//}
	//
	//if s,ok:=p.(string);ok{
	//	fmt.Println("string",s)
	//
	//	return
	//}
	//
	//fmt.Print("unknow")
	switch v := p.(type) {
	case int:
		fmt.Println("integer", v)
	case string:
		fmt.Println("string", v)
	case bool:
		fmt.Println("bool", v)
	default:
		fmt.Print("unknow")
	}

}

func TestEmptyInterface(t *testing.T) {
	doSomething(10)
	doSomething("100")
}
