package main

import (
	"fmt"

	mod "github.com/lppgo/test_mod"
	modv2 "github.com/lppgo/test_mod/v2"
)

func main() {
	fmt.Println("-----------我的go module测试------------")
	fmt.Println(mod.Hi("lucas"))
	fmt.Println("--------在同一个项目中可以使用同一个包的不同版本（看做2个包）--------")
	fmt.Println(modv2.Hi("lucas", "es"))
}
