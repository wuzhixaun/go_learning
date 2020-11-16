package _map

import "testing"

/**
	map初始化
 */
func TestInitMap(t *testing.T)  {

	// 方式1
	m1:=map[string]int{"one": 1, "two": 2, "three": 3}
	t.Log(m1)

	// 方式2,初始化空map
	m2:=map[string]int{}
	m2["name"]= 100
	t.Log(m2)

	// 方式3 make
	m3:=make(map[string]int, 10)
	t.Log(m3,len(m3))

	if v,ok:=m2["name"];ok{
		t.Log(v)
	}

}

// map遍历
func TestTravelMap(t *testing.T)  {
 	m1:=map[string]int{"1": 1, "2": 2, "3": 3, "10": 10}
 	for k,v:=range m1 {
 		t.Log(k, v)

	}

}