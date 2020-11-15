package condition

import "testing"

func TestSwitch1(t *testing.T)  {


	a:= "wuzhixuan"

	switch a {

	case "sy":
		t.Log(a)
	case "wuzhixuan":
		t.Log(a)
	}
}

func TestSwitch2(t *testing.T)  {

	num:= 4


	switch  {
	case 0<num &&num< 5:

		t.Log("0~3")
	case 4<num:
		t.Log(4)



	}
}

func TestMultiCase(t *testing.T)  {

	for i:=0;i<10;i++ {

		switch i {

		case 1,3,5, 7:
			t.Log("even")

		case 2,4,6, 8:
			t.Log("odd")

		default:
			t.Log("is default")
		}
	}

	
}
