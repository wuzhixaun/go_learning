package series

import "fmt"

// 大写的函数可以在包外面访问
func GetFibonacciSerie(n int) []int {

	ret := []int{1, 1}

	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}

	return ret
}

// 小写字母开头的函数不能被包外访问
func square(num int) int {
	return num * num
}

// 执行main方法，所依赖包的所有init方法都会被执行，且只会被执行一次，可以理解为java的静态文件
func init() {
	fmt.Println("init")
}

func init() {
	fmt.Println("init2")
}
