

package main

import "fmt"
import "unicode/utf8"
import "strings"


func string_length() {
	s2 := "世界"
	fmt.Printf("byte length:%d\n", len(s2)) // 字节长度
	fmt.Printf("Rune count:%d\n",utf8.RuneCountInString(s2)) //字符个数
}

func string_split(){
	s1 := "a,b,c"
	result1 := strings.Split(s1, ",")
    fmt.Printf("%q\n", result1)

	s2 := " xyz "
	result2 := strings.Split(s2, "")
    fmt.Printf("%q\n", result2)
}

func strings_splitN(){
	s1 := "a,b,c"
    fmt.Printf("%q\n", strings.SplitN(s1, ",", 2))
}

func decode_utf8_rune(){
    name := "中国行"
    runValue, width:= utf8.DecodeRuneInString(name)
	fmt.Printf("%q, %d", runValue, width)
}

func main(){
	string_length()
	string_split()
	strings_splitN()
    decode_utf8_rune()
}
