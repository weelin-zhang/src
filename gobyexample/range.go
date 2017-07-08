package main

import (
	"fmt"
)

func main() {
	nums:=[]int{2,3,4}
	sum:=0
	for _,i:=range nums{
		sum+=i
	}
	fmt.Println(sum)

	for index,i:=range nums{
		if i == 3 {
			fmt.Println(index)
		}
	}

	kvs:=map[string]string{"a": "apple", "b": "banana"}
	kvs["c"]="add"
	fmt.Println(kvs,kvs["a"],kvs["c"])

	for k,v := range kvs{
		fmt.Println(k,v)
	}

	for k:=range kvs{
		fmt.Println(k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

