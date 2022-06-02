package main

import (
	"fmt"
)

type Number interface {
	int | int32 | int64 | float32 | float64
}

// min int型の引数同士を比較し、大きい方を返す
func min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func main() {
	// ジェネリックを使ってスマートにかこう！
	var a int32 = 10
	var b int32 = 20
	c := min(a, b)
	fmt.Printf("%v\n", c)
}
