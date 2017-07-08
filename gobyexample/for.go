package main

import "fmt"

func main() {
	i:=10
	fmt.Println("i=",i,"&i=",&i)
	for i:=0;i<=3;i++{	//i=0可以覆盖i=10,for中声明的变量作用域只在for中
		fmt.Println("i=",i,"&i=",&i)
	}
	fmt.Println("i=",i,"&i=",&i)

	i1:=1
	fmt.Println("i1=",i1,"&i1=",&i1)
	for i1<=3{
		fmt.Println("i1=",i1,"&i1=",&i1)
		i1++
	}
	fmt.Println("i1=",i1,"&i1=",&i1)

	for{
		fmt.Println(i,i1)
		goto LABEL1
	}
	LABEL1:
	fmt.Println("end1")

	LABEL2:
	for{
		fmt.Println(i,i1)
		break LABEL2
	}

	fmt.Println("end2")

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}