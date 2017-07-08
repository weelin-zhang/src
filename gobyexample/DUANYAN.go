package main

import "fmt"
func ServeHTTP(string) {
	fmt.Println("这他妈的仅仅是一函数,ca")
}

type Handler func(string)

func panduan(in interface{}) {
	if v,ok:=in.(Handler);ok{
		v("wujunbin")
	}
}

func main() {
	panduan(Handler(ServeHTTP))
}
