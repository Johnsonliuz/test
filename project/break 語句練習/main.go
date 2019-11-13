

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	
	var a int
	var b int
	var c string
	var d string
	
	//隨機產生兩個數 當和大於100則跳出
	
	for{
	rand.Seed(time.Now().UnixNano())
	a = rand.Intn(100) + 1
	b = rand.Intn(100) + 1
	
	fmt.Printf("a=%v\n",a)
	fmt.Printf("b=%v\n",b)
	fmt.Printf("總和=%d\n",a+b)
	if a + b > 100{
		break
		}
	}
	
	
	//寫一個登入驗證，有三次機會，帳號為張三豐，密碼為888，正確則提示登入成功，錯誤則提示還有幾次機會
	
	
	
	for i :=1; i <= 3; i++{
	
	fmt.Println("請輸入帳號：")
	fmt.Scanf("%s",&c)
	fmt.Println("請輸入密碼：")
	fmt.Scanf("%s",&d)
	
		if c == "張三豐" && d == "888" {
			fmt.Println("登入成功")
			break
		} else {
			fmt.Printf("登入失敗，還有%v次機會\n",3-i)
		}
	}
}