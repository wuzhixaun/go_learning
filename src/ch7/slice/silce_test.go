package slice

import "testing"

/**
len 指的是切片中的初始化为0的元素，没有被初始化的不能访问
 */
func TestSliceInit(t *testing.T)  {

	// 初始化切片
	var s0 []int
	t.Log(len(s0),cap(s0))

	// 添加元素
	s0 = append(s0, 1)
	t.Log(len(s0),cap(s0))

	// 初始化方式2
	s1 := []int{1, 2, 3, 4, 6}
	t.Log(len(s1),cap(s1))

	// 初始化方式3()
	s2 := make([]int, 3, 5)
	t.Log(len(s2),cap(s2))
	t.Log(s2[0],s2[1],s2[2])

	s2 = append(s2,100)
	t.Log(s2[0],s2[1],s2[2],s2[3])
}


// slice分片
func TestSliceGrowing(t *testing.T)  {
	s:=[]int{}

	for i:=0;i<10;i++ {
		s = append(s, i)
		t.Log(len(s),cap(s)) // 容量是x2的速度
	}
}


// 分片共享存储空
func TestSliceShareMemory(t *testing.T)  {

}
