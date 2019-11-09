//練習
//打印出空心金字塔，外圍用*號包圍

/*思路分析
先打印3*3九宮格塔
接著變成等腰直角三角型如下圖
*
**
***
接著在變成
  *
 ***
*****
最後在變成空心金字塔
*/




//解答-------------













package main

import "fmt"

fnuc main() {
	var level in = 3
	for i := 1; i <= level; i++{
		for a := 1; a <= level-i; a++{
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++{
			if j == 1 || j == 2*i-1 || i == level{
			fmt.Print("*")
			} else{
			fmt.Print(" ")
			}
		}
		fmt.Println("\n")
	}
}