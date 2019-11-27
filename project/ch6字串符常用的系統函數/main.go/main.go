package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	
	//統計字串符長度 len(str)
	str := "Johnson"
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
	fmt.Printf("是否有在seafood裡面：%v\n",str4)
	
	//統計字串裡出現過多少次子串 strings.Count("cheese","ee") //1
	num := strings.Count("cheese","ee")
	fmt.Printf("num=%v\n",num)
	
	//不分大小寫判斷字母是否一樣 fmt.Println(string.EqualFold("abc","ABc")) //true
	b := strings.EqualFold("abc","ABc")
	fmt.Printf("是否一樣:%v\n",b)
	
	fmt.Printf("abc是否等於ABc:%v\n","abc" == "ABc")
	
	//返回子串在字符第一次出現的index值，如果沒有則返回-1 strings.Index("NLT_abcabcabc","abc")
	index := strings.Index("NLT_abcabcabc","abc")
	fmt.Printf("index=%v\n",index)
	
	//返回子串在字符最後一次出現的index值，如果沒有則返回-1 strings.LastIndex("go golang","go")
	index = strings.LastIndex("go golang","go")
	fmt.Printf("index=%v\n",index)
	
	//將指定的子串替換成另外一個子串 strings.Replace("go go hello","嗨嗨",n)n可以指定替換幾個，-1則全部替換
	str2 = "go go hello"
	str = strings.Replace(str2,"go","嗨嗨",-1)
	fmt.Printf("str=%v, str2=%v\n",str,str2)
	
	//按照指定的某個字符，為分割標示，將一個字符拆分成字符串數組 strings.Split("hello,wrold,ok",",")
	strArr := strings.Split("hello,world,ok",",")
	for i := 0; i < len(strArr); i++{
		fmt.Printf("str[%v]=%v\n",i,strArr[i])
	}
	fmt.Printf("strArr=%v\n",strArr)
	
	//將字符串的字母進行大小寫轉換 strings.ToLower("Go")// go strings.ToUpper("Go")//GO
	str = "goLang Hello"
	str = strings.ToLower(str)
	fmt.Printf("str轉變結果:%v\n",str)
	str = strings.ToUpper(str)
	fmt.Printf("str轉變結果：%v\n",str)
	
	//將字符串的左右兩邊的空格去掉 strings.TrimSpace(" tn a lone gopher ntrn ")
	str = strings.TrimSpace(" tn a long gopher ntrn ")
	fmt.Printf("str:%v\n",str)
	
	//將字符串左右兩邊的指定字符去掉 strings.Trim("!hello!","!")
	str = strings.Trim("!hello!","!")
	fmt.Printf("str:%v\n",str)
	
	//將字符串左邊的指定字符去掉 strings.TrimLeft("!hello!","!")
	str = strings.TrimLeft("!hello!","!")
	fmt.Printf("str:%v\n",str)
	
	//將字符串右邊的指定字符去掉 strings.TrimRight("!hello!","!")
	str = strings.TrimRight("!hello!","!")
	fmt.Printf("str:%v\n",str)
	
	//判斷字符串是否以指定的字符串開頭 strings.HasPrefix("ftp://192.168.0.1","ftp") //true
	b =  strings.HasPrefix("ftp://192.168.0.1","ftp")
	fmt.Printf("b=%v\n",b)
	
	//判斷字符串是否以指定的字符串結尾 strings.HasSuffix("NLT_abc.jpg","abc") //false
}