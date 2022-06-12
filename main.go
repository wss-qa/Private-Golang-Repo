package main

import (
	"fmt"
	"github.com/smallnest/ringbuffer"
	"golang.org/x/text/transform"
	"github.com/unknwon/cae/zip"
)

func main() {
	fmt.Println("Hello!!!")
	var buff ringbuffer.RingBuffer
	buff.WriteString("hey")
	fmt.Println(buff.ReadByte())
	n := transform.Chain()
	fmt.Println(n)
	_,_ = zip.Open("bbb")
}
