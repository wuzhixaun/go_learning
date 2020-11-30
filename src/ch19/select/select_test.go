package _select

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 500)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("task is done")
}
func AsyncService() chan string {

	//retCh:= make(chan string)
	retCh := make(chan string, 1)

	go func() {
		ret := service()

		fmt.Println("return result")

		// changle放置
		retCh <- ret

		fmt.Println("service exit")
	}()

	return retCh
}

func TestSelect(t *testing.T) {

	// 超时机制
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Log("time out")
	}

}
