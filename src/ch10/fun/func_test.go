package fun

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 返回两个值
func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func TestFn(t *testing.T) {
	values, i := returnMultiValues()
	t.Log(values, i)

	funcTime := calFuncTime(slowFun)

	funcTime(10)

}

// 计算时长
func calFuncTime(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println(time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

// 求和
func sum(ops ...int) int {

	s := 0

	for _, v := range ops {
		s += v
	}

	return s
}

func TestVarParams(t *testing.T) {
	i := sum(1, 1, 23, 4, 5, 6)
	t.Log(i)
}

// defer 函数，延迟执行函数,函数返回前执行，类似finalize

func TestDefer(t *testing.T) {
	defer func() {
		t.Log("释放资源")
	}()

	t.Log("hello")

	t.Log("函数执行最后")
}
