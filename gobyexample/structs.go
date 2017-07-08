package main

import (
	"fmt"
)

type person struct {
	name string
	age int
}

func main()  {
	a:=person{name:"zhangweijian",age:19}
	b:=person{"zhangweijian",10}
	fmt.Println(a,b)
	fmt.Println(person{"Bob", 20})
	t:=person{"zwj",15}
	fmt.Println(t.age,t.name)
	t_copy:=t
	t_copy.age=100
	fmt.Println(t_copy,t)

	t_point:=&t
	t_point.age=200
	fmt.Println(t_point,t)

}