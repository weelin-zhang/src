package main

import (
	"fmt"
	//"math"
)

const s string="constant"
//var a=10   //常量右边不能出现变量
//const a=a  //错误
const a=10
const con  = a //正确
func main() {
	fmt.Println(con)
	fmt.Println(s)
	const n=10000
	const m  = n/89
	var p  = n/89.0
	const m2 int32=900
	fmt.Println(m)
	fmt.Println(string(m))
	fmt.Println(p)
	fmt.Println(int64(p))
	fmt.Println(float32(m2))

}