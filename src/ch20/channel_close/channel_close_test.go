package channel_close

import (
	"fmt"
	"sync"
	"testing"
)

func dataProducer(ch chan int, wg *sync.WaitGroup) {

	go func() {
		for i := 0; i < 10000; i++ {
			ch <- i
		}

		close(ch)
		wg.Done()
	}()

}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for true {
			if value, flag := <-ch; flag {
				fmt.Println(value)
			} else {

				fmt.Println("channel close")
				break
			}
		}

		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T) {

	var wg sync.WaitGroup

	ch := make(chan int)

	wg.Add(1)

	dataProducer(ch, &wg)

	wg.Add(1)

	dataReceiver(ch, &wg)

	wg.Add(1)

	dataReceiver(ch, &wg)

	wg.Wait()

}
