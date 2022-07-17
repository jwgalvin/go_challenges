package main

import (
	"fmt"
	"strings"
)

func main(){
	str := "hello.person.hooman."
	fmt.Println(ReplaceDots(str))
	}
	// below worked for Kata
func ReplaceDots(str string) string {
	if strings.Contains(str, "." ){
		return strings.Replace(str, ".", "-", -1)
	}else{
		return str
	}
}