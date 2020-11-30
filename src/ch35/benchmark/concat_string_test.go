package benchmark

import (
	"fmt"
	"testing"
)

func BenchmarkConcatStringByAdd(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	fmt.Println("hhhhhhh")
	b.ResetTimer()
	ret := ""

	for _, elem := range elems {
		ret += elem
	}

	b.StopTimer()

	fmt.Println("kkkkkk")
}
