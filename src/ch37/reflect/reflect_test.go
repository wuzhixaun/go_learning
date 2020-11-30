package reflect

import (
	"reflect"
	"testing"
)

type Employee struct {
	EmployeeId string
	Name       string `format:"normal"`
	Age        int
}

func (e *Employee) UpdateAge(newAge int) {
	e.Age = newAge
}

type Customer struct {
	CookieId string
	Name     string
	Age      int
}

func TestInvokeByName(t *testing.T) {

	e := &Employee{"1", "tom", 20}

	// 按名字获取成员
	t.Logf("Name: value(%[1]v),Type(%[1]T)", reflect.ValueOf(*e).FieldByName("Name"), reflect.TypeOf(e))

	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); ok {
		t.Log(nameField)
	} else {
		t.Logf("failed to get name")
	}

	// 重新赋值
	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(1)})

	t.Log(e)

}
