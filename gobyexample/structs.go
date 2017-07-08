package main

import (
	"fmt"
)

type person struct {
	name string
	age int
}

type car struct {
	name string
	price int
	wheel struct{
		num int
	}
}


type teacher struct {
	person
	class string
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


	pp:=&person{name:"nice",age:10}
	age_add(pp)
	fmt.Println(pp.age,pp.name)

	pp2:=person{name:"nice",age:10}
	age_add2(pp2)
	fmt.Println(pp2.age,pp2.name)

	m:= &struct {
		Name string
		Age int
	}{
		Name:"liubi",
		Age:2,
	}

	fmt.Println(m,m.Name,m.Age)


	c:=&car{name:"BMW",price:2000000}
	c.wheel.num=4

	fmt.Println(c)



	tech:=&teacher{class:"English",person:person{name:"li",age:35}}
	fmt.Println(tech,tech.age,tech.name,tech.class)
}

func age_add(per *person)  {
	per.age+=19
	fmt.Println(per.age)
}

func age_add2(per person)  {
	per.age+=19
	fmt.Println(per.age)
}