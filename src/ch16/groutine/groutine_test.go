package groutine

import (
	"fmt"
	"testing"
)

// 协程
func TestGroutine(t *testing.T) {

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)

	}
}
