package main

import "fmt"

// min int型の引数同士を比較し、大きい方を返す
func min(a int64, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
func main() {
	// ジェネリックを使ってスマートにかこう！
	var a int32 = 1
	var b int32 = 2
	c := min(int64(a), int64(b))
	fmt.Printf("%v\n", c)
}
