package csp

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
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

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()

	otherTask()

	fmt.Println(<-retCh)
}
