//寫出一個可以自行輸入兩次數值之後做加減乘除

package main

import (
	"fmt"
	"goproject/ch6/1/utils"
)
	
func main() {
	var a float64
	var b float64
	var res float64
	var operator string
	
	fmt.Println("請輸入第一個數：")
	fmt.Scanln(&a)
	fmt.Println("請輸入運算符：")
	fmt.Scanln(&operator)
	fmt.Println("請輸入第二個數：")
	fmt.Scanln(&b)
		
	res = utils.Cal(a,b,operator)
	fmt.Printf("運算結果=%v\n",res)
}