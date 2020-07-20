package main //声明main包,表明当前是一个可执行程序
import "fmt" //引入内置的fmt包
var (
	person string // ''
	age    int    // 0
	isOk   bool   // false
)
func foo(number int,name string)(int,string){
  return number,name
}
func main(){
	first,_ :=foo(10,"tom")
	_,second :=foo(20,"lily")
	fmt.Print(first)
	fmt.Print(second)
}
