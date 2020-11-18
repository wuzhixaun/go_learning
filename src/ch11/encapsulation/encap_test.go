package encapsulation

import "fmt"

// 创建model
type Employee struct {
	Id   string
	Name string
	Age  int
}

func (e Employee) String() string {
	return fmt.Sprintf("IDL%s-name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func TestCreateEp() {

}
