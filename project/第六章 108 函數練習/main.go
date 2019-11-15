//寫出一個可以自行輸入兩次數值之後做加減乘除

package main

import "fmt"

func cal(a float64,b float64,operator string)(res float64){
		switch operator{
			case "+":
			res = a + b
			case "-":
			res = a - b
			case "*":
			res = a * b
			case "/":
			res = a / b
			default:
			fmt.Println("輸入錯誤")
		}
		return res
	}
	
	
func main() {
	var a float64
	var b float64
	var res float64
	var operator string
	
	fmt.Println("請輸入第一個數：")
	fmt.Scanf("%v",&a)
	fmt.Println("請輸入運算符：")
	fmt.Scanf("%s",&operator)
	fmt.Println("請輸入第二個數：")
	fmt.Scanf("%v",&b)
		
	res = cal(a,b,operator)
	fmt.Printf("運算結果=%v\n",res)
}