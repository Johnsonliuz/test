package main

import "fmt"

func main() {
	var a int
	var b int
	var c int
	var d int
	var e string
	var f int
	
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
	   fmt.Println("輸出結果為B，因為m=0，有進判斷輸出 但n不等於2，故輸出B")
	
	fmt.Println("------------------")
	
	
	//保存用戶名和密碼，判斷用戶是否為張三，密碼是否為1234，如果是則提示登入成功，不是則提示登入失敗
	
	fmt.Println("請輸入一個用戶名：")
	fmt.Scanf("%d",&e)
	fmt.Println("請輸入一個密碼：")
	fmt.Scanf("%d",&f)
	
	if e == '張三' {
		if f == 1234 {
			fmt.Println("登入成功")
		} else {
			fmt.Println("登入失敗")
		}
	}
	
	fmt.Println("------------------")
	
	
	//請輸入一個月份，判斷出當月有幾天（switch）
	
	fmt.Println("請輸入一個月份：")
	fmt.Scanf("%d",&a)
	
	switch a {
		case 1, 3, 5, 7, 8, 10, 12 :
		fmt.Println("31天")
		case 4, 6, 9, 11 :
		fmt.Println("30天")
		case 2 :
		fmt.Println("28/29天")
	}
	
	fmt.Println("------------------")
	
	
	//開發一款軟件，根據提示(身高-108)*2=體重，可以有10公斤的浮動，觀察體重是否合適
	
	fmt.Println("請輸入一個身高：")
	fmt.Scanf("%d",&a)
	fmt.Println("請輸入一個體重：")
	fmt.Scanf("%d",&b）
	
	if (a - 108) * 2 <= b + 5 && (a - 108) * 2 => b - 5{
		fmt.Println("體重合適")
	} else if (a - 108) * 2 < b - 5 {
		fmt.Println("體重過輕")
	} eles if (a - 108) * 2 > b + 5 {
		fmt.Println("體重過重")
	}
	
	fmt.Println("------------------")
	
	
	//輸入一個成績，判斷區間，90-100優秀，80-89優良，70-79良好，60-69及格，60以下不及格
	
	fmt.Println("請輸入一個成績：")
	fmt.Scanf("%d",&a)
	
	switch a {
		case a => 90 && a <= 100:
		fmt.Println("優秀")
		case a => 80 && a <= 89:
		fmt.Println("優良")
		case a => 70 && a <= 79:
		fmt.Println("良好")
		case a => 60 && a <= 69:
		fmt.Println("及格")
		case a < 60 && a => 0:
		fmt.Println("不及格")
		case a < 0:
		fmt.Println("輸入錯誤")
	}
	
	fmt.Println("------------------")
	
	
	//有a,b兩個數，如果a能夠被b整除或a+b大於1000，則輸出a，否則則輸出b
	
	fmt.Println("請輸入第一個整數：")
	fmt.Scanf("%d",&a)
	fmt.Println("請輸入第二個整數：")
	fmt.Scanf("%d",&b)
	
	if a % b == 0 || a + b > 1000{
		fmt.Println("a")
	} else {
		fmt.Println("b")
	}
	
	fmt.Println("------------------")
	
	
	//實現輸入三個整數自動進行排序，由小到大輸出
	
	fmt.Println("請輸入第一個整數：")
	fmt.Scanf("%d",&a)
	fmt.Println("請輸入第二個整數：")
	fmt.Scanf("%d",&b)
	fmt.Println("請輸入第三個整數：")
	fmt.Scanf("%d",&c)
	
	if a < b {
		if a < c {
			if b < c {
				fmt.Println("%d,%d,%d",a,b,c)
			} else {
				fmt.Println("%d,%d,%d",a,c,b)
			}
		} else {
			fmt.Println("%d,%d,%d",c,a,b)
		}
	} else if b < c {
		if a < c {
			if a < b {
				fmt.Println("%d,%d,%d",a,b,c)
			} else {
				fmt.Println("%d,%d,%d",b,a,c)
		} else {
		fmt.Println("%d,%d,%d",b,c,a)
	} else if c < a {
		if b < a {
			if c < b {
				fmt.Println("%d,%d,%d",c,b,a)
			} else {
				fmt.Println("%d,%d,%d",b,c,a)
			}
		} else {
			fmt.Println("%d,%d,%d",c,a,b)
		}
	}
	
	fmt.Println("------------------")
}