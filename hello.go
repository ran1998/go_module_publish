package go_module_publish

import (
	"fmt"
	"github.com/ran1998/go_module_publish/child_package"
)

func Hello() {
	fmt.Println("hello!")
	child_package.Hello()
}
