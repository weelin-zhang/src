package main

import (
	"fmt"
)

type OS interface {
	start_os()
}

type winos struct {
	version string
}


type linuxos struct {
	version string
}
func (os winos)start_os()   {
	fmt.Println("started",os.version)
}


func (os linuxos)start_os()  {
	fmt.Println("starte",os.version)
}

func func_interface(arg interface{}) {

	switch v:=arg.(type) {
	case winos:

		fmt.Println( v.version,"running...")

	case linuxos:
		fmt.Println( v.version,"running...")
	default:
		fmt.Println("Unknown OS")
	}
}

func main(){
	var t1=winos{version:"win10"}
	var t2=winos{version:"win7"}
	var t3=linuxos{version:"centos7"}
	var t4=linuxos{version:"centos6"}
	t1.start_os()
	t2.start_os()
	t3.start_os()
	t4.start_os()
	func_interface(t1)
	func_interface(t2)
	func_interface(t3)
	func_interface(t4)

	aa:=1
	func_interface(aa)
}


