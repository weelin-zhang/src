package main

import (
	"fmt"
	"time"
)

func am_pm(t int)  {
	switch{
	case t>12:
		fmt.Println("afternoon")
	default:
		fmt.Println("morning")
	}
}
func main() {

	var i=2
	fmt.Println("write",i,"as:")
	switch i{
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	fmt.Println(time.Thursday)
	fmt.Println(time.Now().Weekday())
	fmt.Println("today is:")
	switch time.Now().Weekday() {
	case 1:
		println("Mon")
	case 2:
		fmt.Println("Thu")
	case 3:
		fmt.Println("Wen")
	case 4:
		fmt.Println("Thur")
	case 5:
		fmt.Println("Fri")
	case 6:
		fmt.Println("Sat")
	default:
		fmt.Println("Fun")
	}

	var t=time.Now()
	fmt.Println(t.Year(),t.Month(),t.Day(),t.Hour(),t.Minute(),t.Second())
	am_pm(t.Hour())


	add_func:=func(x,y int) int{
		return x+y
	}
	fmt.Println(add_func(10,99))


	whatAmI := func(i interface{}) {
	switch t := i.(type) {
	case bool:
	fmt.Println("I'm a bool")
	case int:
	fmt.Println("I'm an int")
	default:
	fmt.Printf("Don't know type %T\n", t)
	}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
