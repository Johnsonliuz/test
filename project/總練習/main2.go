package main

import(
	"fmt"
	"strconv"
	"strings"
)

fnuc main(){

/*面試題
有兩個變量
a,b
­要求將其進行交換
但不允許使用其中變量
求打印結果*/

	var a int = 3
	var b int = 5

	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("原本的a=3,新a=%v,原本的b=5,新b=%v",a,b)


//練習
	var str string = "ture"
	var c bool
	var str2 string = "1234590"
	var n1 int64
	var n2 int
	var str3 string = "123.456"
	var n3 float64 
	//請用strconv轉換回各自型態
	c , _ = strconv.ParseBool(str)
	n1, _ = strconv.Atoi(str2)
	n2 = int(n1)
	n3, _ = strconv.ParseFloat(str3,64)
	fmt.Printf("c=%T,n1=%T,n2=%T,n3=%T",c,n1,n2,n3)



//練習
	var n4 int = 99
	var n5 float64 = 23.456
	var d bool = true
	//用strconv轉換成str型態
	str = strconv.Itoa(n4)
	str2 = strconv.FormatFloat(n5)
	str3 = strconv.FormatBool(d)
	fmt.Printf("str=%T,str2=%T,str3=%T"str,str2,str3)



//練習
	var str4 string = "ture"
	var e bool
	var str5 string = "1234590"
	var n6 int64
	var n7 int
	var str6 string = "123.456"
	var n8 float64 
	//請用strconv轉換回各自型態
	e, _ = strconv.ParseBool(str4)
	n6, _ = strconv.ParseInt(str5,10)
	n7 = int(n6)
	n8, _ = strconv.ParseFloat(str6,64)
	fmt.Printf("e=%T,n6=%T,n7=%T,n8=%T",e,n6,n7,n8)



//練習
	var n9 int = 99
	var n10 float64 = 23.456
	var f bool = true
	//用strconv轉換成str型態
	str = strconv.Itoa(n9)
	str2 = strconv.FormatFolat(n10)
	str3 = strconv.FormatBool(f)
	fmt.Printf("str=%T,str2=%T,str3=%T",str,str2,str3)


//練習
	var num1 int = 99
	var num2 float64 = 23.456
	var num3 bool = true
	var mychar byte = 'h'
	//將其轉成string型態
	str = strconv.Itoa(num1)
	str2 = strconv.FormatFloat(num2)
	str3 = strconv.FormatBool(num3)
	str4 = strconv.FormatFolat(mychar)
	fmt.Printf("str=%T,str2=%T,str3=%T,str4=%T",str,str2,str3,str4)


//指針練習----------
	var g int = 300
	var ptr *int = &g //錯在哪？如何修正


	var h int = 300
	var ptr float32 = &h //錯在哪？如何修正


	var i int = 300
	var j int = 400
	var ptr2 *int = &i 
	*ptr2 = 100 //i=100
	ptr2 = &j
	*ptr2 = 200 //j=200
	fmt.Printf(i=%d,j=%d,ptr2=%d,i,j,ptr2)
	//輸出什麼內容


//if練習
//編寫一個程序，可以自己輸入年齡
//如果大於18歲則輸出"你的年齡大於18 要對自己的行為負責"
//不到18歲則輸出"多練幾年再來"

	fmt.Println("請輸入年齡：")
	fmt.Scanln(&a)
	if a >= 18{
		fmt.Println("你的年齡大於18,要對自己的行為負責")
	} else if a <18 {
		fmt.Println("多練幾年再來")
	}


//switch練習
//請寫出a-e用小寫轉換成大寫
//其他值則出現錯誤訊息
	var z byte
	fmt.Println("請輸入一個a-e的字母：")
	fmt.Scanln(&z)
	switch {
		case 'a':
		fmt.Println("A")
		case 'b':
		fmt.Println("B")
		case 'c':
		fmt.Println("C")
		case 'd':
		fmt.Println("D")
		case 'e':
		fmt.Println("E")
	}



//進階練習
//學生成績大於等於60分輸出合格，低於60輸出不合格，成績不能高於100分
	fmt.Println("請輸入成績：")
	fmt.Scanln(&a)
	switch (a/60){
		case 1 :
		fmt.Println("及格")
		case 0 :
		fmt.Println("不及格")
	}


//練習
//根據用戶指定的月份，打印該月份所屬的季節
//3.4.5 春
//6.7.8 夏
//9.10.11 秋
//12.1.2 冬

	fmt.Println("請輸入月份：")
	fmt.Scanln(&a)
	switch{
		case 3,4,5:
		fmt.Println("春")
		case 6,7,8:
		fmt.Println("夏")
		case 9,10,11:
		fmt.Println("秋")
		case 12,1,2:
		fmt.Println("冬")
	
	}


//練習
//用for循環寫出"你好,冠陞"


//練習
//請打印出1~100所有9的倍數的整數各數及總和


//練習
//試著寫出0 + 6 = 6 , 1 + 5 = 6依序往下打印出來,6是可變的

//練習
//有三個班，每班有五個學生，求每個班的平均分數及總平均分數
//學生成績用手動輸入
//統計三個班及格人數


//練習
//用*號打印出金字塔


//練習
//打印出空心金字塔，外圍用*號包圍


//打印出九九乘法表


//隨機產生兩個數 當和大於100則跳出


//實現判斷一個整數屬於哪個範圍 >0,<0,=0


//判斷一個年份是否為閏年


//判斷一個整數是否為水仙花數
// 153 = 100*1+10*5+1*3
// 153 = 1*1*1+5*5*5+3*3*3


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


//練習
//從1-100隨機取值，取到99後停止


//從鍵盤輸入幾個不確定的整數，判斷他是正數，負數，遇到O則跳出結束(運用for+continue+break)


//某人有100000元，過路費每次要收5%，少於五萬則每次收1000，請問可以過幾次



}