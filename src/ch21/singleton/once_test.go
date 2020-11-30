package singleton

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {
}

var singleInstance *Singleton

var once sync.Once

func GetSingletonInstance() *Singleton {

	once.Do(func() {
		fmt.Println("GetSingletonInstance")
		singleInstance = new(Singleton)

	})

	return singleInstance
}

func TestSingleton(t *testing.T) {

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			instance := GetSingletonInstance()
			t.Log(unsafe.Pointer(instance))
		}()
	}
	wg.Wait()
}
