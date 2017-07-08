package main
import "fmt"

type rect struct {
width, height int
}


func (r *rect) area() int {   //这里拿到的是r的地址，假如修改r的属性，r的属性会改变
	fmt.Println("执行了吗")
	fmt.Println("666")
return r.width * r.height
}



func (r rect) perim() int {  //这里拿到的是r的拷贝，假如修改r的属性，r的真实的属性不会改变
	fmt.Println("执行了吗")
	fmt.Println("666")
return 2*r.width + 2*r.height
}



func main() {
r := rect{width: 10, height: 5} //width看可以省去



fmt.Println("area: ", r.area())
fmt.Println("perim:", r.perim())


//
//rp := &r
//fmt.Println("area: ", rp.area())
//fmt.Println("perim:", rp.perim())


	var tz1,tz2 TZ
	tz1,tz2=10,10

	tz1.print()
	tz2.print2()

	fmt.Println(tz1,tz2)
	fmt.Printf("tz1:%p,tz2:%p\n",&tz1,&tz2)

	var tz3 TZ
	fmt.Printf("%p",&tz3)
	fmt.Println(tz3)
	tz3.Increase()
	fmt.Println(tz3)



}
//其它类型也有方法
type TZ int

func (z TZ) print(){
	fmt.Println(z,"并不是只有结构类型有方法")
	fmt.Printf("print:%p\n",&z)
	//尝试修改值
	z=11

}

func (z *TZ) print2(){
	fmt.Println(z,"并不是只有结构类型有方法")
	fmt.Printf("print1:%p\n",z)
	*z=12
}

func (z *TZ) Increase(){
	*z=100
}