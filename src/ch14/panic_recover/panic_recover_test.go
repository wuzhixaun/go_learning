package panic_recover

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T) {

	fmt.Println("Start")

	//os.Exit(10)

	defer func() {

		fmt.Println("defer")
	}()

	panic(errors.New("调用栈信息"))
}
