package operate_test

import "testing"


// 简单运算符
func TestOperate(t *testing.T)  {

	// 定义数组
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 3, 4}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 3, 2, 4}
	e := [...]int{1, 2, 3, 4}

	t.Log(a== b)
	// t.Log(a== c)  // 数组个数不同的不能进行比较
	t.Log(a== d)
	t.Log(a== e)
}

