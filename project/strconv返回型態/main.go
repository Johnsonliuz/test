�m��
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
    //�Х�strconv�ഫ�^�U�۫��A
}


�ѵ�
package main

import (
    "fmt"
    "strconv"
)

func main() {
    var str string = "ture"
    var b bool
    var str2 string = "1234590"
    var n1 int64
    var n2 int    //�n����int����
    var str3 string = "123.456"
    var n3 float64 
    //�Х�strconv�ഫ�^�U�۫��A

    b, _ = strconv.ParseBool(str)
    fmt.Printf("b=%t b���A=%T\n",b,b)
    n1, _ = strconv.ParseInt(str2,10,64)
    n2 = int(n1)
    fmt.Printf("n2=%v n2���A=%T\n",n2,n2)
    n3, _ = strconv.ParseFloat(str3,64)
    fmt.Printf("n3=%f n3���A=%T\n",n3,n3) 
}