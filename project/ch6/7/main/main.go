//編寫一個函數 makeSuffix，可以接受一個帶有後綴肚的值，返回一個閉包
//在閉包內判斷是否有後綴，若沒有後綴則加上後綴
//strings.HasSuffix 此函數可以判斷是否有後綴
package main

import (
	"fmt"
	"strings"
)

func makeSuffix(suffix string) func(name string) string {
	
	return func(name string) string{
	if !strings.HasSuffix(name,suffix){
		return name + suffix
	}
	
	return name
	}
}

//如果不使用閉包如何完成？
func makeSuffix2(suffix string,name string)string{
	
	if !strings.HasSuffix(name,suffix){
		return name + suffix
	}
	return name

}

func main() {
	
	a := makeSuffix(".jpg")
	fmt.Println("name=",a("aaa"))
	b := makeSuffix2(".avi","bbb")
	fmt.Println("name=",b)
}