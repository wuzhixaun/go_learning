package array

import "testing"

func TestArrayInit(t *testing.T) {
	// 定义长度为4的数组
	var arr [4]int

	// 方式二
	arr1 := [5]int{1, 2, 3, 45}

	t.Log("arr1", arr1)

	arr[0] = 100

	t.Log(arr, arr[0], arr[1], arr[2])

	// 方式三
	arr3 := [...]int{1, 2, 3, 4, 5, 6, 7, 89, 10}

	t.Log("arr3", arr3)

}

// 数组遍历
func TestArrayLoop(t *testing.T) {

	// 初始化数组
	arr3 := [...]int{1, 2, 3, 4, 5, 6, 7, 89, 10}

	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}
}
