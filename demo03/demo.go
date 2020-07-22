package main
import "fmt"
import "strings"
func main(){
	// fmt.Print("hello")
	// fmt.Print("我们把go语言称为\"golang\"")
	// var str string = `
	// 弃我去者,昨日之日不可留
	// 乱我心者,今日之日多烦忧`
	// fmt.Print(str)
	// fmt.Print(len(str))
	// say := "hello"
	// hi := "world"
	// sayhi := say+hi
	// fmt.Print(say+hi)
	// fmt.Print(sayhi)
	// var str string = "123456"
	// fmt.Print(strings.HasPrefix(str,"123"))
	var str = "c:\\windows\\system\\user\\index"
	var str2 = strings.Split(str,"\\")
	fmt.Print(strings.Join(str2,"+"))
}