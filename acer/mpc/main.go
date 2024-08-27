package main

import (
	"fmt"
//	"sync"
)

type Bal struct {
	msg string
}

func (b *Bal) sendMsg(msg string) {
	if b.msg == "" {
		b.msg = msg
	}
}

func (b *Bal) readMsg() string {
	if b.msg != "" {
		return b.msg
	}
}


func main() {
  var bal Bal
  bal.msg=""
  bal.sendMsg("hello")
  fmt.Println(bal.readMsg())
}
