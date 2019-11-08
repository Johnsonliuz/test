//練習
package main

import "fmt"

func main() {
    var str string = "ture"
    var b bool
    var str2 string = "1234590"
    var n1 int64
    var n2 int
    var str3 string = "123.456"
    var n3 float64 
    //請用strconv轉換回各自型態
}


//解答--------------
















/*
package main

import (
    "fmt"
    "strconv"
)

func main() {
    var str string = "true"
    var b bool
    var str2 string = "1234590"
    var n1 int64
    var n2 int    //­要換成int型態
    var str3 string = "123.456"
    var n3 float64 
    //請用strconv轉換回各自型態

    b, _ = strconv.ParseBool(str)
    fmt.Printf("b=%t b型態=%T\n",b,b)
    n1, _ = strconv.ParseInt(str2,10,64)
    n2 = int(n1)
    fmt.Printf("n2=%v n2型態=%T\n",n2,n2)
    n3, _ = strconv.ParseFloat(str3,64)
    fmt.Printf("n3=%f n3型態=%T\n",n3,n3) 
}