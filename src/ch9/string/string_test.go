package string

import "testing"

func TestString(t *testing.T) {
	var str string
	t.Log(str) // string 类型初始化值为""

	str = "hello"
	t.Log(len(str))

	str = "\xE4\xB8\xA5" // 可以存储任何二进制数据
	t.Log(str)
	t.Log(len(str)) // string 是不可变的byte slice

	str = "中"

	// rune 取出字符串的unicode
	c:=[]rune(str)
	t.Log(len(c))
	t.Logf("中的 unicode %x",c[0])
	t.Logf("中的 utf-8%x",str)

	//
}

func TestStringToRune(t *testing.T)  {
	s := "中华人民共和国"

	for _,c:=range s{
		t.Logf("%[1]c,%[1]d", c)
	}
}
