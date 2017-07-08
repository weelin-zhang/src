package main
import "fmt"

type rect struct {
width, height int
}


func (r *rect) area() int {
	fmt.Println("执行了吗")
	fmt.Println("666")
return r.width * r.height
}



func (r rect) perim() int {
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
}

