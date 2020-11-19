package encapsulation

import (
	"fmt"
	"testing"
)

// 创建model
type Employee struct {
	Id   string
	Name string
	Age  int
}

func (e Employee) String() string {
	return fmt.Sprintf("ID:%s-name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func TestCreateEp(t *testing.T) {

	e := &Employee{"0", "Bob", 20}

	t.Log(e.String())
}
