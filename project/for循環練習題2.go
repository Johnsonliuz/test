//練習
//試著寫出0 + 6 = 6 , 1 + 5 = 6依序往下打印出來,6是可變的

//思路分析 要有兩個變數 還有一個最大值

package main

import "fmt"

func main() {
	
	var a int = 6
	for i:= 0; i <= a; i++{
	fmt.Printf("%v + %v = %v\n",i ,a-i ,i+a-i)
	}
}