//輸入位置n可以判斷出費波那契數相對應的數字

package main

import "fmt"

func add(n int)(a int){
	if n == 1 {
		a = 1
	} else if n >= 2 {
		a = add(n - 1) + add(n - 2)
	}
	return a
}

func main() {
	var a int
	var b int
	fmt.Println("請輸入一個位置：")
	fmt.Scanln(&a)
	b = add(a)
	fmt.Printf("值=%v",b)
}