package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	
	//統計字串符長度 len(str)
	str := Johnson
	fmt.Println("str長度：",len(str))
	
	//切片字串符內含有中文出現亂碼 r:=[]rune(str)
	str2 := "哈囉強森"
	r :=[]rune(str2)
	for i := 0; i < len(r); i++{
	fmt.Printf("字符:%c\n",r[i])
	}
	
	//字串符轉整數 n,err := strconv.Atoi("12")
	n,err := strconv.Atoi("12")
	if err != nil{
		fmt.Println("轉換錯誤",err)
	} else {
		fmt.Println("轉換結果是",n)
	}
	
	//整數轉字符串 str := strconv.Itoa(12345)
	str3 := strconv.Itoa(12345)
	fmt.Printf("轉換結果是=%s,型態=%T\n",str3,str3)
	
	//字串符轉byte var bytes = []byte("hello")
	var bytes = []byte("hello")
	fmt.Printf("轉換結果是=%c,型態=%T\n",bytes)
	
	//byte轉字符串 str = string([]byte{97,98,99})
	str = string([]byte{97,98,99})
	fmt.Printf("str=%s\n",str)
	
	//10進制轉2.8.16進制 str = strconv.FormatInt(132,2)// 2->8,16
	str = strconv.FormatInt(132,2)
	fmt.Printf("對132做2進制結果是：%v\n",str)
	str = strconv.FormatInt(132,16)
	fmt.Printf("對132做16進制結果是：%v\n",str)
	
	//找查子串是否有在字串之中 strings.Contains("seafood","sea") //true
	str4 := strings.Contains("seafood","sea") //true
	fmt.Printf("是否有在seafood裡面：%v",str4)
	
	//
	
}