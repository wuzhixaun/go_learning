package remote_package

import (
	cm "github.com/easierway/concurrent_map" // 给报取了一个别名
	"testing"
)

func TestConcurrentMap(t *testing.T) {

	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))

}
