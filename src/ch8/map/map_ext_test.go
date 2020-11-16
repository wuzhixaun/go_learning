package _map

import "testing"


// map，value 作为一个函数
func TestMapWithFunValue(t *testing.T)  {

	m1 := map[int]func(op int) int{}

	m1[0] = test1
	m1[1] = test2
	m1[2] = test3

	t.Log(m1[0](2), m1[1](2), m1[2](2))
}


// 使用map构造Set
// key type 数据类型
// value bool
func TestMapForSet(t *testing.T)  {
	mySet := map[int] bool{}

	// 添加元素
	mySet[1] = true
	mySet[2] = true

	t.Log(mySet[2])

	t.Log(mySet[3])

	delete(mySet, 1)
	t.Log(mySet[1])
}


func test1(a int) int {
	return a
}

func test2(a int) int {
	return a * a
}

func test3(a int) int {
	return a * a * a
}


