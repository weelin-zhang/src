package main

import "fmt"

func main() {
	m:=make(map [string]int)

	fmt.Println(m)

	m["k1"]=7
	m["k2"]=13
	fmt.Println("map:",m)

	v1:=m["k1"]
	fmt.Println("v1:",v1)
	fmt.Println(len(m))

	delete(m,"k1")
	fmt.Println(m)
	_,prs:=m["k32"]				//判断是否存在"k32"这个key
	fmt.Println("prs:",prs)
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
