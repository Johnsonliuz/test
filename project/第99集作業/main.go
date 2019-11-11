package main

import "fmt"

func main() {
	var a int
	var b int
	var c int
	var d int
	
	//實現判斷一個整數屬於哪個範圍 >0,<0,=0
	
	fmt.Println("請輸入一個整數：")
	fmt.Scanf("%d",&a)
	if a > 0 {
		fmt.Println("正數")
	} else if a < 0{
		fmt.Println("負數")
	} else if a == 0 {
		fmt.Println("0")
	}
	
	fmt.Println("------------------")
	
	
	//判斷一個年份是否為閏年
	
	fmt.Println("請輸入一個年份：")
	fmt.Scanf("%d",&a)
	
	if a % 4 == 0 {
		fmt.Println("%d是閏年",a)
	} else {
		fmt.Println("%d不是閏年",a)
	}
	
	fmt.Println("------------------")
	
	
	//判斷一個整數是否為水仙花數
	// 153 = 100*1+10*5+1*3
	// 153 = 1*1*1+5*5*5+3*3*3
	
	fmt.Println("請輸入一個整數：")
	fmt.Scanf("%d",&a)
	b = a / 100
	c = (a - 100*b) / 10
	d = (a - 100*b - 10*c) / 1 
	
	if a = b*b*b+c*c*c+d*d*d && a = 100*b+10*c+1*d {
		fmt.Println("是水仙花數")
	} else {
		fmt.Println("不是水仙花數")
	}
	
	fmt.Println("------------------")
	
	
	//寫出輸出結果
	/* m := 0
	   n := 3
	   if m = 0 {
	   if n = 2 {
	   fmt.Println("A")
	   } else {
	   fmt.Println("B")
	   }
	   }
	   */
	   
	//保存用戶名和密碼，判斷用戶是否為張三，密碼是否為1234，如果是則提示登入成功，不是則提示登入失敗
	
	//請輸入一個月份，判斷出當月有幾天（switch）
	
	//開發一款軟件，根據提示(身高-108)*2=體重，可以有10公斤的浮動，觀察體重是否合適
	
	//輸入一個成績，判斷區間，90-100優秀，80-89優良，70-79良好，60-69及格，60以下不及格
	
	//有a,b兩個數，如果a能夠被b整除或a+b大於1000，則輸出a，否則則輸出b
	
	//實現輸入三個整數自動進行排序，由小到大輸出
}