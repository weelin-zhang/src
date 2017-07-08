package main
import "errors"
import (
	"fmt"
	"reflect"
)

//By convention, errors are the last return value and have type error, a built-in interface.


func f1(arg int) (int, error) {
	if arg == 42 {

		//errors.New constructs a basic error value with the given error message.


		return -1, errors.New("can't work with 42")



	}

	//A nil value in the error position indicates that there was no error.


	return arg + 3, nil
}

//It’s possible to use custom types as errors by implementing the Error() method on them. Here’s a variant on the example above that uses a custom type to explicitly represent an argument error.


type argError struct {
	arg  int
	prob string
}

var testargError argError=argError{100,"nice"}


func (e argError) Error() string {			//定义了一个函数
	fmt.Println("老子Error被调用了..")
	//return fmt.Sprintf("%d - %s", e.arg, e.prob)
	return "40 - can't work with it"
}



func f2(arg int) (int, error) {
	if arg == 42 {

		//In this case we use &argError syntax to build a new struct, supplying values for the two fields arg and prob.

fmt.Println("确定什么时候调用了Error")
return -1, argError{arg, "can't work with it"}
}
return arg + 3, nil
}



func main() {

	//The two loops below test out each of our error-returning functions. Note that the use of an inline error check on the if line is a common idiom in Go code.


	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	//If you want to programmatically use the data in a custom error, you’ll need to get the error as an instance of the custom error type via type assertion.


	fmt.Println("===========")
	_, e := f2(42)
	fmt.Println("e的类型是:",reflect.TypeOf(e)," ,e=",e)
	fmt.Println("--------------")

	if ae, ok := e.(argError); ok {     //什么意思？？？？
		//fmt.Println(ok,"ae=",ae)
		fmt.Println("xa:",ae.arg)
		fmt.Println("xb:",ae.prob)
	}

	fmt.Println("自定义一个argError类型的变量")
	fmt.Println(testargError.arg,testargError.prob)
	fmt.Println(testargError.Error())
}
