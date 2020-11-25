package share_mem

import (
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {

	count := 0

	for i := 0; i < 100; i++ {

		go func() {
			count++
		}()

	}

	t.Log(count)
}

func TestCounterSafe(t *testing.T) {
	var mut sync.Mutex
	count := 0
	for i := 0; i < 10; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			count++
		}()

	}
	time.Sleep(1 * time.Second)
	t.Log(count)
}

func TestCounterWaitGroup(t *testing.T) {
	var mut sync.WaitGroup

	count := 0
	for i := 0; i < 10; i++ {
		mut.Add(1)

		go func() {
			defer func() {

			}()
			count++
			mut.Done()
		}()
	}

	mut.Wait()
	t.Log(count)
}
