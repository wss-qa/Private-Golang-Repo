package hello

import (	
	"fmt"	
	"github.com/unknwon/cae/zip"
)

func Hello() {
	fmt.Println("Hello, World!")
	_,_ = zip.Open("aaa")
}
