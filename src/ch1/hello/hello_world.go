package main   // 包

import (
	"fmt"
	"os"
) // 导入代码依赖

// 功能实现
func main()  {

	if len(os.Args) > 1 {
		fmt.Print(os.Args)
	}else {
		fmt.Print("c参数值长度小于1")
	}


	fmt.Println("Hello world")

	// 返回退出状态
	os.Exit(100)
}