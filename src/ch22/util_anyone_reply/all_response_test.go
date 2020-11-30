package util_anyone_reply

import "testing"

func AllResponse() string {

	num := 10

	// 创建一个管道
	ch := make(chan string)

	for i := 0; i < num; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)

	}

	finalRet := ""

	for i := 0; i < num; i++ {
		finalRet += <-ch + "\n"

	}
	return finalRet
}

func TestAllFirstResponse(t *testing.T) {

	t.Log(AllResponse())
}
