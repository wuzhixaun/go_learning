package constant

import "testing"


const (
	Monday = iota + 1
	Tu
	We
) // 常量定义
func TestConstant(t *testing.T)  {

	t.Log(Monday,Tu, We)



}
