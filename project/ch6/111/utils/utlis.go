package utils

import "fmt"

func Cal(a float64,b float64,operator string)(res float64){
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