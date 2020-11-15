package _type

import (
	"testing"
)

type Myint int64


// 类型转换，没有隐式类型转换
func TestTypeTran(t *testing.T)  {


	var a int = 1

	var b int64

	b = int64(a)

	var c Myint

	c = Myint(b)

	t.Log(c)

	t.Log(b)
}


func TestPoint(t *testing.T)  {
	a:=1

	// 利用取地址符号获取地址
	aptr := &a
	t.Log(a, aptr)

	// 格式化T获取变量类型
	t.Logf("%T %T",a, aptr)

}


/**
	字符串
 */
func TestString(t *testing.T)  {
	var str string

	t.Log("*"+str+ "*")
	t.Log(len(str))
}