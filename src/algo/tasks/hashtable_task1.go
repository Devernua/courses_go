package main

import (
	"fmt"
	"algo/hashtable"
)

type notePhone struct {
	number string
	name string
}

func (n notePhone) Data() (result []byte) {
	for _, c := range n.number {
		result = append(result, byte(c))
	}
	return result
}

func main() {
	var n int
	fmt.Scanf("%v", &n)

	h := hashtable.New(uint32(n) + 1)

	for i := 0; i < n; i++ {
		var cmd, number string
		fmt.Scanf("%v", &cmd)
		fmt.Scanf("%v", &number)

		if cmd == "add" {
			var name string
			fmt.Scanf("%v", &name)
			h.Insert(&notePhone{number, name})
		} else if cmd == "del" {
			h.Delete(&notePhone{number, ""})
		} else if cmd == "find" {
			res := h.Find(&notePhone{number, ""})
			if res != nil {
				fmt.Println(res.(*notePhone).name)
			} else {
				fmt.Println("not found")
			}
		} else {
			panic("invalid command")
		}
	}
}
