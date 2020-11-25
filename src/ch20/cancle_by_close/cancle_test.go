package cancle_by_close

import (
	"fmt"
	"testing"
	"time"
)

func TestCancleTask(t *testing.T) {

	cancleChan := make(chan struct{}, 0)

	for i := 0; i < 5; i++ {
		go func(i int, cancleChan chan struct{}) {
			for {
				if isCancled(cancleChan) {
					break
				}
				time.Sleep(time.Microsecond * 5)
			}

			fmt.Println(i, "cancelled")
		}(i, cancleChan)
	}

	cancle_2(cancleChan)
	time.Sleep(time.Second * 1)

}

func isCancled(cancelChan chan struct{}) bool {

	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

// 往changle放置消息
func cancle_1(cancleChan chan struct{}) {
	cancleChan <- struct{}{}
}

func cancle_2(cancleChan chan struct{}) {
	close(cancleChan)
}
