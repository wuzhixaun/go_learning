package loop

import "testing"

func TestWhileLoop(t *testing.T)  {
	i:= 0


	for i< 5 {
		t.Log(i)
		i++
	}


	for {
		t.Log(i)
		i++
	}

}
