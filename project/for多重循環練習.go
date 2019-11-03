//練習
//用*號打印出金字塔

package main

import "fmt"

func main() {

	var max int = 5
	
	for j := 1; j <= max; j++{
		for a := 1; a <= max-i; a++{
		fmt.Print(" ")
		}
		for i := 1; i <= 2*j-1; i++{
		fmt.Print("*")
		}
	fmt.Println()
	}
}