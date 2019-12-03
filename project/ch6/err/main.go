//用一個函數去讀取文件名稱init.conf的信息
//如果文件名不正確則返回一個自定義錯誤

package main

import (
	"fmt"
	"errors"
)

func a(name string)(err error){
	if name == "init.conf"{
		return nil
	} else {
		return errors.New("讀取文件錯誤") //返回自定義錯誤
	}
}

func b() {
	err := a("aaa")
	if err != nil{
		panic(err) //如果文件錯誤則跳出
	}
	fmt.Println("正常執行")
}

func main() {
	b()
	fmt.Println("AAAA")
}

