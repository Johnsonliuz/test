//練習
//請寫出a-e用小寫轉換成大寫
//其他值則出現錯誤訊息

package main

import "fmt"

func main() {
    var char byte
    
    fmt.Println("請輸入一個字母")
    fmt.Scanf("%c", &char)
    
    switch char {
        case 'a' :
        fmt.Println(A)
        case 'b' :
        fmt.Println(B)
        case 'c' :
        fmt.Println(C)
        case 'd' :
        fmt.Println(D)
        case 'e' :
        fmt.Println(E)
    }
}