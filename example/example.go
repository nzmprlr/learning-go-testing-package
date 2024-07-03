package example

import "fmt"

func PrintMap(x int) {
	m := map[int]int{}
	for v := range x {
		m[v] = v
	}

	for v := range m {
		fmt.Println(v)
	}
}
