package main

import (
	"fmt"

	"github.com/exercise/src"
)

func main() {
	list := []int64{}
	for i := 0; i < src.Length; i++ {
		list = append(list, int64(i))
	}
	val := src.ConcurencyHTTP(list, 10)
	fmt.Println(len(val))
}
