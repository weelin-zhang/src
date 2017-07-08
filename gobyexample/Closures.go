package main

import "fmt"

func closuretest(x int) func(int)int{
	return func(y int)int{
		x++
		fmt.Printf("y==%p\n",&y)
		fmt.Printf("x==%p\n",&x)
		return x
	}
}




func main() {
	v:=10
	fmt.Printf("v==%p\n",&v)
	f:=closuretest(v)
	fmt.Println(f(1))
	fmt.Println(f(1))
	fmt.Println(f(1))

}
