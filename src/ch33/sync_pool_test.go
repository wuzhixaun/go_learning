package ch33

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {

	pool := sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new Project")
			return 100
		},
	}

	v := pool.Get().(int)

	t.Log(v)

	pool.Put(18)

	v1 := pool.Get().(int)

	t.Log(v1)
	v2 := pool.Get().(int)
	t.Log(v2)

	pool.Put(28)

	runtime.GC() // GC  会清除sync.pool 缓存的对象

	v3 := pool.Get().(int)
	t.Log(v3)

}
