package main

import (
	"fmt"
	"algo/hashtable"
)

type stringWrapper struct {
	Str string
}

func (str *stringWrapper) Data() (buf []byte) {
	// TODO: optimize
	for _, c := range str.Str {
		buf = append(buf, byte(c))
	}
	return buf
}


func main() {
	var m uint32
	var n int
	fmt.Scanf("%v", &m)
	fmt.Scanf("%v", &n)

	h := hashtable.New(m)

	for i := 0; i < n; i++ {
		var cmd string
		fmt.Scanf("%v", &cmd)

		if cmd == "add" {
			var val string;
			fmt.Scanf("%v", &val)
			h.Insert(&stringWrapper{val})
		} else if cmd == "del" {
			var val string;
			fmt.Scanf("%v", &val)
			h.Delete(&stringWrapper{val})
		} else if cmd == "find" {
			var val string;
			fmt.Scanf("%v", &val)
			res := h.Find(&stringWrapper{val})
			if res != nil {
				fmt.Println("yes")
			} else {
				fmt.Println("no")
			}
		} else if cmd == "check" {
			var idx uint32
			fmt.Scanf("%v", &idx)
			res := h.Check(idx)
			for idx, val := range res {
				fmt.Print(val.(*stringWrapper).Str)
				if idx + 1 != len(res) {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		} else {
			panic("invalid command")
		}
	}
}
