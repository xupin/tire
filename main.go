package main

import (
	"fmt"

	"github.com/xupin/tire/tire"
)

func main() {
	node := &tire.Node{
		Char: "/",
		Next: make(map[rune]*tire.Node, 0),
	}
	tire := &tire.Tire{
		Root: node,
	}
	tire.Append("123")
	tire.Append("12")
	tire.Remove("123")
	fmt.Printf("共有[%d]个字符 \n", tire.Size)
	if tire.Find("12", true) {
		fmt.Println("能查到")
	} else {
		fmt.Println("不能查到")
	}
}
