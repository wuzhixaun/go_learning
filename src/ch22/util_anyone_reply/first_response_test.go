package util_anyone_reply

import (
	"fmt"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(time.Millisecond * 10)
	return fmt.Sprintf("the result is from %d", id)
}

func FirstResponse() string {

	num := 10

	// 创建一个管道
	ch := make(chan string)

	for i := 0; i < num; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)

	}
	return <-ch
}

func TestFirstResponse(t *testing.T) {

	t.Log(FirstResponse())
}
