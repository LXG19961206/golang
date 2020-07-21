package main
import "fmt"
import "math"
const pi int = 123
func main(){
	const (
	  a = iota
	  b = 100
	  c 
	  d = iota
	  e 
	)
	// fmt.Print(a,b,c,d,e)
	// test1()
	test2()
	// test3()
	test4()
}
func test1(){
	const (
		a,b = iota,iota
		c,d = iota,iota
	)
	fmt.Print(a,b,c,d)
}
func test2(){
	var a int =10
  fmt.Printf("%d \n",a)  //10
  fmt.Printf("%b \n",a)  //1010  占位符%b 表示的是2进制的数
  //8进制
  var b int = 077
  fmt.Printf("%o \n",b)  //77
  fmt.Printf("%d\n",b)
  //16进制
  var c int = 0xff
  fmt.Printf("%x \n",c) //ff 
  fmt.Printf("%d \n",c) //FF 
  var d int = 0x1234
  fmt.Printf("%T\n",d)
}

func test3(){
  fmt.Printf("%f\n",math.Pi)
  fmt.Printf("%2f\n",math.Pi)
  f1 := 1.2345
  fmt.Printf("%T\n",f1) //go语言默认小数都是float64类型
  f2 :=float32(1.23456)
  fmt.Printf("%T\n",f2) //显示声明float32类型
}

func test4(){
	a := true
	var b bool 
	fmt.Printf("%T\n",a)
	fmt.Print(b)
}