package main

import "fmt"

func add(x,y int)(m,n int)  {
	m=y
	n=x
	//return
	return m+n,m-n
}

func add2(x,y int)(int,int)  {
	return x+y,x-y
	//return 				Error error

}

func main() {
	fmt.Println(add(1,9))
	fmt.Println(add2(1,9))

	sum(1,2)

	sum(1,2,3)

	nums:=[]int{1,2,3,4}
	sum(nums...)
	sum(1,2,3,4)

}

func sum(nums...int){
	fmt.Println(nums," ")
	total:=0
	for _,num:=range nums{
		total+=num
	}
	fmt.Println(total)
}