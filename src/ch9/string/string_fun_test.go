package string

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T)  {

	s :="A,B,C,C,D,E"
	arr := strings.Split(s, ",")

	t.Log(arr)

	// 遍历输出(id就是下标值，如果想不输出则可以用_代替)
	for id,value:=range arr {
		t.Log(id, value)
	}

	// 连接
	str :=strings.Join(arr,"＋")

	t.Log(str)
}

func TestStringConv(t *testing.T)  {

	// int转换成str
	str1 := strconv.Itoa(5)
	t.Log(str1)




	if i,err:=strconv.Atoi("8");err == nil {
		t.Log(i, err)

	}

}
