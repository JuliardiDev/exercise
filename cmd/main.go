package main

import (
	"exercise/src"
	"fmt"
)

func main() {
	list := []int64{}
	for i := 0; i < src.Length; i++ {
		list = append(list, int64(i))
	}

	fmt.Println(src.ConcurencyHTTP(list, 10))
}
