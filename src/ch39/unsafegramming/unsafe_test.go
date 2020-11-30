package unsafegramming

import (
	"testing"
	"unsafe"
)

func TestUnsafe(t *testing.T) {

	i := 10
	f := *(*float64)(unsafe.Pointer(&i))

	t.Log(f)
}

type myInt int

func TestMyInt(t *testing.T) {

	i := 1000
	i2 := (*myInt)(unsafe.Pointer(&i))

	t.Log(i2)
}
