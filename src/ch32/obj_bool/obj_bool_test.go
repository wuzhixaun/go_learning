package obj_bool

import (
	"fmt"
	"testing"
	"time"
)

func TestObjectPool(t *testing.T) {
	// 初始化一个对象池

	objPool := NewObjPool(10)

	for i := 0; i < 12; i++ {
		if v, err := objPool.GetObRes(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			if err := objPool.ReleaseObject(v); err != nil {
				t.Error(err)
			}
		}
	}

	fmt.Println("done")
}
