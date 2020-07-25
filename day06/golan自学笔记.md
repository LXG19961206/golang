#### golan自学笔记

###### golan ？

* 编译型语言 

  * 与解释型的区别

    > 解释型语言: 人类可读的代码 => 虚拟机(解释器) => 处理器

    > 编译型语言: 人类可读的代码 => 处理器

  * 编译型语言不需要一个中间的解释器,运行起来性能更高

* 21世纪的`c语言`

  > 1.支持交叉编译,编译快速
  >
  > 2.开发效率高 (自带垃圾回收机制)
  >
  > 3.执行性能好 
  >
  > 4.能最大限度地发挥出硬件的性能(天生支持并发)
  >
  > 5.语法简洁
  >
  > 6.企业级的编程语言

###### 开发环境搭建

* 编译器 : 

  * go语言采用的是`utf-8`编码

  * 可以使用`vscode`,因为免费
    * 下载安装包
    * 一路下一步
    * 安装插件`chinese`,`go`
  * 也可以使用官方编译器,但是比较贵

* 安装开发环境包

  ```go
  https://golang.google.cn/dl/
  ```

* 查看版本

  ```go
  go version
  ```

* 配置GOPATH

  ```shell
   1 export GOROOT=/usr/local/go
   2 export GOPATH=/Users/apple/Desktop/golearn
   3 export PATH=$PATH:${GOPATH//://bin:}/bin
  ```

  ```jsx
  go env    //查看go环境配置
  ```

  

  

  

###### 第一个go程序

* 创建一个`demo.go`文件

  ```go
  package main //声明main包,表明当前是一个可执行程序
  package xxx  //如果不是main包,那就不会生成一份可执行文件
  import "fmt" //引入内置的fmt包
  func main(){
    // 主函数,是程序的执行入口,go程序默认从main入口函数开始执行
    fmt.Println("hello world")
  }
  ```

* 编译

  * 使用 `go build`

    > 1.cd到 go文件所在的目录
    >
    > 2.使用 go build
    >
    > 3.会将go文件进行编译

  * 如果是在其他路径里,写入该文件在 `src`往后的路径,生成可执行文件在你打开终端的路径

    ```shell
    go build github.com/LXG19961206/demo01
    ```

  * 如果想要规定输入的文件的文件名

    ```
    go build -o demo
    ```

  * 使用 `go run`

    > 1.像执行脚本文件一样执行go代码

  * 使用`go install`

    > 1.先编译得到一个可执行文件
    >
    > 2.会将这个可执行文件拷贝到 `GOPATH\bin`目录

  * 如果有跨平台编译(交叉编译),请详见相关文档教程



* 变量的声明

  * 在函数外只能放置 标识符 的声明,不能放置一些逻辑语句

    > 何为标识符 ?
    >
    > * 所谓标识符,就是程序员的具有特殊意义的词,比如
    >   * 变量
    >   * 常量
    >   * 函数
    >   * 类型
    >   * ...
    > * go语言中标识符由字母和下划线组成,并且只能以字母或者`_`开头,比如
    >   * abc
    >   * _
    >   * _123
    >   * a123

  * golan中的关键字

    | break    | default     | func   | interface | select |
    | -------- | ----------- | ------ | --------- | ------ |
    | case     | defer       | go     | map       | struct |
    | chan     | else        | goto   | package   | switch |
    | const    | fallthrough | if     | range     | type   |
    | continue | for         | import | return    | var    |


###### 变量和常量

* 变量

  * 为何出现变量

    > 程序运行需要的数据都保存在内存里,但是如果我们通过内存地址操作这个数据的话,代码可读性很差
    >
    > 因此出现了变量,方便我们通过变量去找到内存里对应的地址

  * go中的变量

    * 需要使用`var`关键字,先声明然后再去使用

    * 是一种静态语言,声明时必须标明类型

      > string
      >
      > int
      >
      > bool
      >
      > ...

      ```go
      var person string
      var age int
      ```

    * 可以使用批量声明法

      ```go
      var (
      	person string
      	age int
      )
      ```

    * 声明了变量如果没有赋值,值为对应类型的空值

      ```go
      var (
      	person string // ''
      	age    int    // 0
      	isOk   bool   // false
      )
      ```

    * 变量赋值,直接可以使用`=`就可以进行赋值

      ```go
      var (
      	person string // ''
      	age    int    // 0
      	isOk   bool   // false
      )
      func main(){
        person = "tom"
        fmt.Println("hello"+" "+person)
      }
      ```

      * 可以在声明 的同时进行赋值

        ```go
        var (
        	person string = "tom"
        	age int = 23
        )
        ```

        

      * 几种打印方式

        ```go
        func main(){
          fmt.print("hello world")  //直接打印出指定的内容
        }
        ```

        ```go
        func main(){
        	fmt.Printf("name:%s",person)
        }
        // %s其实就是占位符
        // 之后会采用person的值来替代的这个占位符 
        ```

        ```go
        func main(){
        	fmt.Println(person)  //会在终端中打印出相应的内容后在后面添加一个换行符
          fmt.Prinitln()       //可以用于快速打印出一个空行
        }
        ```

        

    * 但是,在`go`语言之中,在函数中变量的声明了必须要使用,不然无法编译

      * 出于减少无效变量出现的所占用的内存的需求,减少性能消耗
      * 在全局里声明的变量可以不使用

    * go语言推荐使用驼峰式的命名法

      * 小驼峰 firstSecond
      * 大驼峰 FirstSecond

    * 类型推导

      > 会根据值判断变量是什么类型,因此可在声明的时候不去写类型,但是这种方式必须在声明的时候赋值
      >
      > ```go
      > func main(){
      > 	var name = "tom"
      > 	fmt.print(name)
      > }
      > ```

    * 短变量声明

      > 简短变量声明法,只能在函数里面使用
      >
      > ```go
      > func main(){
      > 	name := "tom"     //其实就是 var name = "tom" 的简写,一样必须在声明的时候就赋值
      > 	fmt.Print(name)
      > }
      > ```

    * 匿名变量

      > 在使用多重赋值时候,如果想要忽略某个值,可以使用匿名变量
      >
      > ```go
      > func foo(number int,name string)(int,string){
      >   return number,name  //比如是这种可以一次返回多个值的函数
      > }
      > func main(){
      > 	first,_ :=foo(10,"tom")		//匿名变量可以选择忽略某个值
      > 	_,second :=foo(20,"lily")  
      > 	fmt.Print(first)
      > 	fmt.Print(second)
      > }
      > ```

    * 其他注意

      > * 函数外的每个语句都必须以关键字开始
      >   * var
      >   * const
      >   * func
      >   * ...
      > * := 不能用在函数外
      > * _多用于占位,表示忽略值
      > * 同一个作用域中变量不可以重复声明

* 常量的声明

  * 何为常量 ?

    > 在程序运行中恒定不变的量

  * 特性

    > * 恒定不变
    > * 必须在声明的时候就赋值
    > * 赋值之后不可改变
    > * 不可重复声明

  * 批量声明

    >  ```go
    > const(
    > 	STATUSOK = 200
    > 	notFound = 404
    > )
    >  ```
    >
    > * 在批量声明中,如果某一行没有写值,就代表它的值和上面的是一样的
    >
    >   ```go
    >   const (
    >   	a = 100
    >   	b
    >   	c
    >   )
    >   // 则a、b、c都为100
    >   ```

  * 常量计数器

    > * `iota`是go语言的常量计数器,只能在常量的表达式里使用
    >
    > * 在`const`关键字出现时被重置为0
    >
    > * `const`每次新增一行常量声明都会使`iota`计数一次
    >
    > * 使用`iota`可以简化定义,在定义枚举的时候很有用
    >
    >   ```go
    >   const (
    >   	a = iota  //0
    >   	b					//1
    >   	c         //2
    >   	d         //3
    >   	e         //4
    >   )
    >   ```
    >
    > * 一道面试题
    >
    >   ```go
    >   const (
    >   	a = iota
    >   	b
    >   	_
    >   	c
    >   )
    >   ```
    >
    >   * 分析执行过程
    >     * a = 0
    >     * b = 1
    >     * _ = 2  
    >       * 因为是匿名变量因此被抛弃
    >     * c = 3
    >   * 因此结果是 0 ,1 ,3
    >
    > * 面试题 2
    >
    >   ```go
    >   const (
    >   	a = iota
    >   	b = 100
    >   	c 
    >   	d
    >   )
    >   ```
    >
    >   * 分析执行过程
    >     * a = 0
    >     * b = 100
    >     * c没赋值和上面的一样 为100
    >     * d同理
    >   * 结果是 0 ,100 ,100 ,100
    >
    > * 面试题 3
    >
    >   ```go
    >   const (
    >   	a iota
    >     b 100
    >     c 
    >     d iota
    >     e 
    >   )
    >   ```
    >
    >   * 分析执行过程
    >     * a = 0
    >     * b = 100 , const每次新增一行常量声明,iota都会++,因此iota=1
    >     * c = 100 , 没写就和上面一致 , 同理iota ++ 为2
    >     * d = 3
    >     * e = 4
    >   * 因此结果是 0 ,100 ,100 ,3 , 4
    >
    > * 面试题 4
    >
    >   ```go
    >   cosnt (
    >   	a,b = iota,iota
    >   	c,d = iota,iota
    >   )
    >   ```
    >
    >   * 分析执行过程
    >     * 因为 a和b的声明在一行
    >     * 每声明一行 ,iota才会++
    >     * 因此a,b是相等的,都是0
    >     * 此行执行完毕后iota++,因此c,d都是1
    >
    > * 定义数量级(了解)
    >
    >   ```go
    >   const (
    >   	_  = iota
    >     KB = 1 << (10*iota)
    >     MB = 1 << (10*iota)
    >     GB = 1 << (10*iota)
    >     TB = 1 << (10*iota)
    >     PB = 1 << (10*iota)
    >   )
    >   ```

###### 基本数据类型

* 整形

  * 有符号(负号)

    > int8
    >
    > int16
    >
    > int32
    >
    > Int64

  * 无符号

    > uint8
    >
    > uint16
    >
    > uint32
    >
    > uint64

  * 架构特定(取决于系统位数)

    > uint
    >
    > int

  * 类型别名

    > Unicode中的rune等价于int32
    >
    > byte等价于 uint8

  * 特殊类型

    > uintptr (无符号整形)
    >
    > 由系统决定占用位大小，足够存放指针即可，和C库或者系统接口交互

    取值范围

    | 具体类型 | 取值范围                                  |
    | -------- | ----------------------------------------- |
    | int8     | -128到127                                 |
    | uint8    | 0到255                                    |
    | int16    | -32768到32767                             |
    | uint16   | 0到65535                                  |
    | int32    | -2147483648到2147483647                   |
    | uint32   | 0到4294967295                             |
    | int64    | -9223372036854775808到9223372036854775807 |
    | uint64   | 0到18446744073709551615                   |

* 八进制&十六进制

  ```go
  package main
  import "fmt"
  func main(){
    //10进制
    var a int =10
    fmt.Printf("%d \n",a)  //10    占位符%d 表示的是10进制的数
    fmt.Printf("%b \n",a)  //1010  占位符%b 表示的是2进制的数
    //8进制
    var b int = 077
    fmt.Printf("%o\n",b)  //77     占位符%o 表示的是8进制的数
    fmt.Printf("%d\n",b)  //63     %d转为十进制的数
    //16进制
    var c int = 0xff
    fmt.Printf("%x\n",c) //ff      占位符%x 表示的是16进制的数
    fmt.Printf("%d\n",c) //255     转为10进制的数
  }
  ```

* 如何参看变量的类型

  > ```go
  > var d int = 0x1234
  > fmt.Prinf("%T\n",d)   // int
  > ```

* 浮点型

  > ```go
  > package main
  > import "math"
  > import "fmt"
  > func main(){
  >   fmt.Print("%f\n",math.Pi)
  >   fmt.Print("%2f\n",math.Pi)
  >   f1 := 1.2345
  >   fmt.Printf("%T\n",f1) //go语言默认小数都是float64类型
  >   f2 :=float32(1.23456)
  >   fmt.Printf("%T\n",f2) //显示声明float32类型
  > }
  > ```
  >
  > * 因为float64和float32不是一个数据类型,go语言是一个静态类型的语言,因此两者不能相互赋值

* 复数型(忽略)

* 布尔值

  > 1.布尔类型的变量默认为`false`
  >
  > 2.go语言中不允许将整形强制转为布尔型
  >
  > 3.布尔型无法参与数值运算,也无法和其他的语言进行转换
  >
  > ```go
  > func test4(){
  > 	a := true   
  > 	var b bool 
  > 	fmt.Printf("%T\n",a)  //bool
  > 	fmt.Print(b)          //false
  > }
  > ```

* 字符串

  > * go语言中的字符串必须是双引号包裹的
  >
  > * 单引号只能包裹字符
  >
  >   ```go
  >   str1 := "hello world"
  >   str2 := 'h'   //单独的字母、汉字、符号表示一个字符
  >   ```
  >   
  > * 字符和字节
  >
  >   字节:         一个字节=8bit ( 8个二进制位 )
  >
  >   1个字符:  "A"=1个字节
  >
  >   1个汉字:  一般占3个字节
  >
  > * 常见的转义符
  >
  >   | 转义符 | 含义   |
  >   | ------ | ------ |
  >   | \r     | 回车符 |
  >   | \n     | 换行符 |
  >   | \t     | 制表符 |
  >   | \\'    | 单引号 |
  >   | \\"    | 双引号 |
  >   | \\\    | 反斜杠 |
  >
  >   * 比如我们打印一个路径,因为会默认被认为是转义字符,因此如果想要打出一个反斜杠 ,可以连续打两个
  >
  >     ```go
  >     fmt.Print("c:\\window\\system32\\xxx")
  >     fmt.Print("我们把go语言称为\"golang\"")  //打印双引号同理
  >     ```
  >
  > * 多行的字符串
  >
  >   * 定义多行的字符串需要使用反引号
  >
  >   * 反引号中的内容会原样输出
  >
  >     ```go
  >     var str string = `
  >     	弃我去者,昨日之日不可留
  >     	乱我心者,今日之日多烦忧
  >     `
  >     var path string = `c:\windows\system\xxx`  //反引号的话就无须使用转义了,因为回原样输出
  >     fmt.Print(str)
  >     ```
  >
  > * 获取字符串长度
  >
  >   * 使用`len()`函数获取
  >
  >     ```go
  >     var str string = "hello world"
  >     fmt.Print(len(str))
  >     ```
  >
  > * 字符串的拼接
  >
  >   * 字符串拼接(使用加号)
  >
  >     ```go
  >     say := "hello"
  >     hi := "world"
  >     sayhi := say+hi
  >     fmt.Print(say+hi)
  >     ```
  >
  >   * 字符串拼接(使用%s占位符),就可以使用多个占位符
  >
  >     ```go
  >     say := "hello"
  >     hi := "world"
  >     var str = fmt.Sprintf("%s%s",say,hi)   //Sprintf不会执行打印,而是返回拼接结果
  >     fmt.Println(str)
  >     ```
  >
  > * 字符串分割
  >
  >   * 需要引入strings包
  >
  >     ```go
  >     import "strings"
  >     ```
  >
  >   * 使用 `stings.Splits`来进行分割
  >
  >     * 第一个参数是要分割哪个字符串
  >     * 第二个参数是以何为界进行分割
  >
  >     ```go
  >     var path string = "c:\\windows\\system\\user\\xxx"
  >     var str = strings.Split(path,"\\")
  >     fmt.Println(str)   // [c: windows system user xxx]  返回类似一个列表或者是说是数组
  >     ```
  >
  > * 判断是否包含
  >
  >   * `strings.Contains()`函数
  >
  >     * 第一个参数是要判断哪个字符串
  >     * 第二个参数是包含什么内容
  >     * 最后返回一个bool
  >
  >     ```go
  >     var str string = "123456"
  >     fmt.Print(strings.Contains(str,"1"))
  >     ```
  >
  > * 判断前缀 / 后缀
  >
  >   * 前缀 (是否以...开头)
  >     * 使用`strings.HasPrefix(str,"123")`
  >     * 依旧返回一个boo
  >   * 后缀(是否以...结尾)
  >     * 使用`strings.HasSuffix(str,"123")`
  >     * 依旧返回一个boo
  >
  > * 返回字符串里包含...的位置
  >
  >   * 使用`strings.Index(str,"2")`
  >   * 返回的是该字符(子字符串)出现的位置
  >
  > * 返回最后一处包含...的位置
  >
  >   * 使用`strings.LastIndex(str,"2")`
  >
  > * 将数组或者是列表拼接起来
  >
  >   * 使用`strings.Join(str,"")`函数
  >
  >     ```go
  >     var str = "c:\\windows\\system\\user\\index"
  >     var str2 = strings.Split(str,"\\")
  >     fmt.Print(strings.Join(str2,"+"))
  >     //c:+windows+system+user+index
  >     fmt.Printf("$T,%d\n",str2,str2) // 打印出类型和编码好
  >     ```
  >
  > * 基本类型之间的转换
  >
  >   * go语言是一个静态类型语言,定义/赋值/运算时候类型必须一致
  >
  >   * 但是允许我们去手动转换,但是必须是兼容类型,比如布尔不能转为整形
  >
  >     ```go
  >     var a int8 
  >     a = 10
  >     var b int16
  >     b = int16(a)               // 这一步相当于先将int8类型的a转为int16类型,再把值赋给b
  >     fmt.Printf("%T,%s\n",b,b)
  >     ```
  >
  >     ```go
  >     var a = 3.14
  >     var b int 
  >     b = int(a)   // 浮点型转为整形会直接舍弃小数部分
  >     fmt.Println(b)
  >     ```
  >
  >     ```go
  >     var b = true
  >     a = int8(b)   // 在golang中,布尔类型的值不能转为数值型
  >     ```
  >
  >   * 常数会在有需要的时候自动转换
  >
  >   * 变量需要我们手动转型
  >
  >   



###### 运算符

* 算术运算符

  ```go
  + - * / % ++ --
  ```

  ```go
  func main(){
  	a := 10
  	b := 3
  	sum := a+b
  	fmt.Printf("%d,%d = %d\n",a,b,sum)   // 10+3=13
      sub := a-b
      fmt.Printf("%d,%d = %d\n",a,b,sub)   // 10-3=7
  }
  ```

  * 在go语言中,因为是一个静态类型语言,也就是等号两边的值应该是一个类型的,因此在除法时候有些例外

    * 假如等号左边都是整数
    * 返回值也会是整数,因此会舍弃小数部分

    ```go
    func main(){
    	a := 10
    	b := 3
    	res := a/b
    	fmt.Printf("%d,%d = %d\n",a,b,res)   // 10/3=3
        res2 := a%b
        fmt.Printf("%d,%d = %d\n",a,b,res2)   // 10%3=1
    }
    ```

  * ++ 和 -- 不能和其他语言一样参与运算 

    ```go
    var a int = 2
    a++
    fmt.Println(a)
    ```

    * 出于精简语言的需要

* 关系运算符

  * 关系运算符的返回结果都是布尔值

  ```go
  < > <= >= == !=
  ```

  ```go
  package main
  func main(){
      a := 3
      b := 5
      c := 3
      res1 := a<b           //true
      res2 := a>b           //fasle
      res3 := a==b          //false
      res4 := a!=b          //true
      fmt.Printf("%T,%t\n",res1,res1)           //%t是布尔值的占位符
      fmt.Printf("%T,%t\n",res2,res2)
      fmt.Printf("%T,%t\n",res3,res3)
      fmt.Printf("%T,%t\n",res4,res4)
  }
  ```

* 逻辑运算符

  ```go
  && || !
  ```

  > &&  "与"
  >
  > 只有所有的操作数都为真,返回结果才是真,有一个是假的,结果就是假
  >
  > ```go
  > a := true
  > b := true
  > c := false
  > res := a&&b&&c
  > fmt.Printf("%T,%t\n",res,res)   //false
  > ```

  > ||  "或"
  >
  > 只要有一个是真的,结果就为真,如果所有的操作数都是假,那么结果就为假
  >
  > ```go
  > a := true
  > b := true
  > c := false
  > res := a||b||c
  > fmt.Printf("%T,%t\n",res,res)   //true
  > ```

  > !  "非"
  >
  > 会反转操作数的逻辑状态 
  >
  > ```go
  > a := true
  > b := true
  > fmt.Print(a==b)   //true
  > fmt.Print(a==!b)  //false
  > ```

  > 练习题
  >
  > ```go
  > a := 3
  > b := 2
  > c := 5
  > res := a>b && c%a==b && a<(c/b)  //false
  > ```
  >
  > * go语言的`&&`只要是遇到一个`false`就会停止继续运算
  > * go语言的`||`只要是遇到一个`true`就会停止继续运算

* 位运算符

  * 何为位运算: 将数值转为二进制后,按位操作

  * 按位&

    > 对应位的值如果都为1才为1,有一个为0就为0

  * 按位|

    > 对应位的值如果都是0才为0,有一个为1就为1

  * 异或^

    > 二元: a^b
    >
    > ​	对应的值不同为1,相同为0
    >
    > 一元: ^a
    >
    > ​	按位取反 
    >
    > ​			1--->0
    >
    > ​			0--->1

  * 位清空 &^

    > * 对于 a &^ b
    > * 对于 b 上的每个数值
    >
    > * 如果为0,则取a对应位上的数值
    >
    > * 如果为1,则结果位就取0

  ```go
  a := 60
  b := 13
  // a: 0011 1100
  // b: 0000 1101
  // &  0000 1100
  // |  0011 1101
  // ^  0011 0001
  // &^ 0011 0000
  res1 := a&b
  fmt.Printf("%d\n",res1)
  res2 := a|b
  fmt.Printf("%d\n",res2)
  res3 := a^b
  fmt.Printf("%d\n",res3)
  res4 := a&^b
  fmt.Printf("%d\n",res4)
  ```

  * 按位左移 <<

    > 左边的操作数的值向左移动由右操作数指定的位数
    >
    > * 其实就是把 a 转为二进制,向左移动 n 位 ,其实就是放大 2的n次幂
    >
    >   ```go
    >   var a = 4
    >   var b = a >> 2 // 1
    >   var c = a << 2 // 16
    >   ```
    >
    >   

  * 按位右移 >>

    > 左边的操作数的值向右移动由右操作数指定的位数
    >
    > * 其实就是把 a 转为二进制,向右移动 n 位 ,其实就是缩小 2的n次幂
    > * 代码示例同上

* 赋值运算符

  ```go
  = += -= *= /= %= <<= >>= $= ^= |= ^= ...
  ```

  > a +=b   
  >
  > 其实就是 a = a+b
  >
  > a -=b
  >
  > 其实就是 a= a-b
  >
  > ...以此类推 

###### 键盘输入和打印输出

* fmt打印输入

  > * fmt.Print
  >
  > * fmt.Println
  >
  > * fmt.Printf(" ") ,也叫做格式化输出,也就是可以使用各种占位符
  >
  >   | 占位符 |                   含义                   |
  >   | :----: | :--------------------------------------: |
  >   |   %v   |                 原值输出                 |
  >   |   %%   |               单纯的百分号               |
  >   |   %T   |               输出数据类型               |
  >   |   %t   |                 输出bool                 |
  >   |   %b   |                  二进制                  |
  >   |   %c   |           打印数值所对应的字符           |
  >   |   %d   |                  十进制                  |
  >   |   %o   |                  八进制                  |
  >   |   %x   |                 十六进制                 |
  >   |   %X   |            十六进制  大写A-F             |
  >   |   %s   |                  字符串                  |
  >   |   %q   | 该值对应的单引号括起来的go语法字符字面值 |
  >   |  ...   |                   ...                    |

* 键盘输入

  > * fmt.Scanln(&x)
  >
  >   * 会阻塞程序执行
  >   * 之后我们键盘输入的值 ,会被赋值给x
  >
  >   ```go
  >   var x int 
  >   var y float64
  >   fmt.Println("请输入一个整数类型,一个浮点类型")
  >   fmt.Scanln(&x,&y)
  >   // fmt.Scanf("%d,%f",&x,&y)也可以
  >   ```


###### if分支语句

* 程序的流程控制

  * 顺序结构:从上而下,依次执行
  * 选择结构:满足条件,条件才会执行

  * 循环结构:满足条件,代码会被执行多次

* 语法格式

  > if(bool条件的表达式){
  >
  > ​	代码体,当bool表达式的返回值是true的时候会执行
  >
  > }else{
  >
  > ​	代码体,当bool表达式的返回值是false的时候会执行
  >
  > ​	所以如果if中的代码体没有执行,就回去执行else的代码体
  >
  > ​	如果if中的条件满足了,就不会去执行else中的语句
  >
  > }
  >
  > ```go
  > package main
  > 
  > import "fmt"
  > 
  > func main() {
  > 	num := 20
  > 	if num > 16 {
  > 		fmt.Printf("num大于16")
  > 	} else {
  > 		fmt.Printf("num小于16")
  > 	}
  > }
  > ```
  >
  > ```go
  > func main(){
  > 	var num int 
  >   fmt.Print("请输入一个数字")
  >   fmt.Scanln(&num)  // 会监控读取我们键盘输入行为,把我们输入的数赋值给num
  >   if(num>60){
  >     fmt.Print("及格")
  >   }else{
  >     fmt.Print("没及格")
  >   }
  > }
  > ```
  >
  > if(bool条件表达式1){
  >
  > ​	如果bool表达式1的结果是true,那就执行代码体
  >
  > }else if(bool条件表达式2){
  >
  > ​	如果bool表达式2的结果是true,那就执行代码体
  >
  > }else if(bool表达式3){
  >
  > ​	如果bool表达式3的结果是true,那就执行代码体
  >
  > }else{
  >
  > ​	如果所有的 if 和 else if 都不满足,就会执行最后的else
  >
  > }
  >
  > ```go
  > func test2(){
  > 	var num int 
  > 	fmt.Println("请输入一串数字")
  > 	fmt.Scanln(&num)
  > 	if(num<60){
  > 		fmt.Println("不及格")
  > 	}else if(num<70){
  > 		fmt.Println("良好")
  > 	}else if(num<80){
  > 		fmt.Println("优秀")
  > 	}else if(num<90){
  > 		fmt.Println("good")
  > 	}else{
  > 		fmt.Println("极佳")
  > 	}
  > }
  > ```
  >
  > * 当第一个if没有成立的时候 ,证明整形num >= 60的,因此之后的else if没有考虑到60以下的情况
  > * 同样第二个else if没有考虑到num<70的情况,因为程序如果能够执行到这里,证明其必大于等于70
  > * ...
  > * 依次类推,可以用来精简代码

* if语句其他写法

  > ```go
  > func main(){
  >   if num:=4;num>0{
  >     fmt.Println("正数")
  >     fmt.Println(num)
  >   }else{
  >     fmt.Println("负数")
  >   }
  >   fmt.Println(num)
  > }
  > ```
  >
  > * 在写if语句的时候声明一个变量 
  >   * if分号前是我们声明的变量,分号后是判断的条件
  >   * 程序体是一样的
  > * 此变量只能在if语句里才能访问,也就是作用域只在if语句中
  > * 出了if语句就无从访问

###### switch语句

* 语法

  > * 使用switch后的变量名和每个case的数值一一对比
  >
  > * 如果满足条件,就执行case后的代码体
  >
  > * 如果不满足继续往下找
  >
  > * 如果都不满足,就执行default里的代码体
  >
  >   ```go
  >   func test4() {
  >   	var num int
  >   	fmt.Println("please input a number")
  >   	fmt.Scanln(&num)
  >   	switch num {
  >   	case 6:
  >   		fmt.Println("不及格")
  >   	case 7:
  >   		fmt.Println("及格")
  >   	case 8:
  >   		fmt.Println("还可以")
  >   	case 9:
  >   		fmt.Println("良好")
  >   	case 10:
  >   		fmt.Println("优秀")
  >   	default:
  >   		fmt.Println("未知的分数")
  >   	}
  >   }
  >   ```
  >
  >   * switch可以作用于其他类型上,但是case后的数值类型必须和switch保持一致
  >
  >   * Test 模拟实现了一个计算器
  >
  >     ```go
  >     func test5() {
  >     	var (
  >     		num1 int
  >     		num2 int
  >     		opa  string
  >     	)
  >     	fmt.Println("请输入两个数,一个运算符号")
  >     	fmt.Scanln(&num1, &num2, &opa)
  >     	switch opa {
  >     	case "+":
  >     		fmt.Print(num1 + num2)
  >     	case "-":
  >     		fmt.Print(num1 - num2)
  >     	case "*":
  >     		fmt.Print(num1 * num2)
  >     	case "/":
  >     		fmt.Print(num1 / num2)
  >     	}
  >     }
  >     ```

* switch的其他写法

  > 可以省略省略switch后面的变量,相当于直接作用在true上
  >
  > ```go
  > switch{
  > 	case true:
  > 	 	fmt.Println("true")
  > 	case false:
  > 		fmt.Println("false")
  > }
  > ```
  >
  > 因此可以这样写
  >
  > ```go
  > func test6(){
  >   var num int
  >   fmt.Println("请输入一个分数")
  >   switch{
  >     case num<60:
  >     fmt.Println("没及格")
  >     case num<70:
  >     fmt.Println("刚及格")
  >     case num<80:
  >     fmt.Println("还可以")
  >     case num<90:
  >     fmt.Println("良好")
  >     case num<100:
  >     fmt.Println("优秀")
  >     default:
  >     fmt.Println("满分")
  >   }
  > }
  > ```
  >
  > 此外,case后可以同时写多个数值
  >
  > ```go
  > func test7(){
  > 	var month int
  >   fmt.Println("请输入1-12中的任意一个整数")
  >   fmt.Scanln(&month)
  >   switch(month){
  >     case 1,2,3:
  >     fmt.Println("一季度")
  >     case 4,5,6:
  >     fmt.Println("二季度")
  >     case 7,8,9:
  >     fmt.Println("三季度")
  >     case 10,11,12:
  >     fmt.Println("四季度")
  >   }
  > }
  > ```
  >
  > switch后可以多一个初始化语句,类似if的用法,定义一个只能在switch结构体里用的变量
  >
  > ```go
  > func test8(){
  >   switch lang:=20,lang{
  >    ....  //同理,这个lang变量只能在此switch分支里才可以用
  >   }
  > 	// 出了这个作用域访问lang,就得到undefined
  > }
  > ```

* switch中的break 和 fallthrough

  > break可以使用在switch之中,也可以使用在for循环里,可以强行退出结构体
  >
  > ```go
  > func test9(){
  > 	n:=2
  > 	switch n {
  >     case 1:
  >     fmt.Println("1")
  >     break
  >     fmt.Println("1")
  >     fmt.Println("1")
  >     fmt.Println("1")
  >   	case 2:
  >     fmt.Println(2)
  >     fmt.Println(2)
  >     fmt.Println(2)
  >     fmt.Println(2)
  > 	}
  > }
  > ```
  >
  > * 在每个case语句上 其实有自动补上一个break
  > * 但是如果我们在某处手动加上break的话,可以强制退出switch case结构体

  

  >fallthrough 当满足某个case后不但会执行自己,也会无视条件执行下一个case
  >
  >```go
  >func test10(){
  >	n:=1
  >	switch n {
  >    case 1:
  >    fmt.Println("1")
  >		fallthrough
  >  	case 2:
  >    fmt.Println(2)
  >	}
  >}
  >//结果是 1,2
  >```

###### for循环

* 何为

  > 循环就是在条件满足的前提下,不断执行一段相同或者相似的代码

* 语法

  > * 结构
  >
  >   ```go
  >   func main(){
  >     for 循环变量;循环条件;循环变量的变化{
  >       //循环执行的代码
  >     }
  >   }
  >   ```
  >
  > * 例子,打印五次 helloworld
  >
  >   ```go
  >   func main(){
  >     for i := 1;i<=5;i++{
  >   		fmt.Println("hello world")
  >   	}
  >   }
  >   ```
  >
  >   * 在for循环里的循环变量 i 作用域是当前这个for循环,出了循环外部不能使用
  >   * 只是使用`i:=1`,不能使用`var i=1`,同样在`if`和`switch`语句之中也是这样

* 其他写法

  > * 虽然是有三个循环表达式,但是可根据情况进行省略
  >
  >   * 可以同时省略循环变量声明以及循环变量的变化,只写循环条件
  >
  >     ```go
  >     i:=5
  >     for i<5 {
  >     	fmt.Println("hello world")
  >     }
  >     ```
  >
  >   * 可以省略三个循环条件,循环体代码就会无条件地不断得到执行,相当于其他语言的`while true`
  >
  >     ```go
  >     for {
  >       fmt.Println("hello world") //会无条件地一直执行打印
  >     }
  >     ```
  >
  > * domo
  >
  >   * 打印从 58 - 23 的数字
  >
  >     ```go
  >     for i:=58;i<=23;i--{
  >     	fmt.Println(i)
  >     }
  >     ```
  >
  >   * 累加1-100的和
  >
  >     ```go
  >     var sum int = 0
  >     for i:=1 ; i<=100 ; i++{
  >     	sum += i
  >     }
  >     fmt.Println(sum)
  >     ```
  >
  >   * 打印 1-100 内,能够被3整除,但是不能被 5 整除的数 ,统计被打印的数字的个数, 每行打印五个
  >
  >     ```go
  >     for i, sum := 1, 0; i <= 100; i++ {
  >     		if i%3 == 0 && i%5 != 0 {
  >     			fmt.Print(i,"\t")
  >     			sum++
  >     			if sum%5 == 0 && sum != 0 {
  >     				fmt.Println()
  >     			}
  >     		}
  >     	}
  >     ```
  >
  >     * 使用 := 声明多个变量要这么写  变量1,变量2,变量3 := 值1,值2,值3
  >     * 5个一行,一次只要个数取余5为0的话,就代表已经满一行,因此就可以换行了
  >
  >   * 求 100 - 999 之间所有的水仙花数
  >
  >     ```go
  >     // 也就是 个位 十位 百位 三位数之和 等于其本身的数字
  >     for i := 100; i <= 999; i++ {
  >     		var temp1 int = i / 100
  >     		var temp2 int = (i - (temp1 * 100)) / 10
  >     		var temp3 int = i % 10
  >     		if temp1*temp1*temp1+temp2*temp2*temp2+temp3*temp3*temp3 == i {
  >     			fmt.Printf("%d,是水仙花数", i)
  >     			fmt.Println()
  >     		}
  >     	}
  >     
  >     // 出于简便,我们可以使用math包里的Pow函数 ,不过要将数值转为 float64才可以使用
  >     for i := 100; i <= 999; i++ {
  >     		var temp1 int = i / 100
  >     		var temp2 int = (i - (temp1 * 100)) / 10
  >     		var temp3 int = i % 10
  >     		if math.Pow(float64(temp1), 3)+math.Pow(float64(temp2), 3)+math.Pow(float64(temp3), 3) == float64(i) {
  >     			fmt.Printf("%d,是水仙花数", i)
  >     			fmt.Println()
  >     		}
  >     	}
  >     
  >     // 当然还有一个更容易理解的方法
  >     for x := 1 ; x<=9 ; x++{
  >       for y :=0 ; y<=9 ; y++ {
  >         for z :=0 ;z<=9 ; z++ {
  >           var temp = 100*x + 10*y + z
  >           if(x*x*x + y*y*y + z*z*z == temp){
  >             fmt.Printf("%d,是水仙花数", temp)
  >     				fmt.Println()
  >           }
  >         }
  >       }
  >     }
  >     ```
  >
  >     
  >
  >   * 打印2-100之内所有的素数
  >
  >     ```go
  >     func test8() {
  >     	for i := 2; i <= 100; i++ {
  >     		var isAnswer bool = true
  >     		for j := 2; j <= 100; j++ {
  >     			if j > i {
  >     				break
  >     			}
  >     			if i%j == 0 && i != j {
  >     				isAnswer = false
  >     			}
  >     		}
  >     		if isAnswer == true {
  >     			fmt.Printf("%d是素数", i)
  >     			fmt.Println()
  >     		}
  >     	}
  >     }
  >     ```
  >
  >     * 外层函数每次循环都会有一个变量记录这个数是否是素数
  >     * 内层函数判断
  >       * 如果被除数大于除数,直接break
  >       * 如果这个数是质数,那就把记录是否为素数的变量更改为false
  >     * 如果执行了内层循环,isAnswer还是true的话,证明当前的这个数就是素数
  >
  >     

* 多层for循环

  > * 绘制乘法口诀表
  >
  >   ```go
  >   func main() {
  >   	fmt.Println("开始")
  >   	for i := 1; i <= 9; i++ {
  >   		for j := 1; j <= i; j++ {
  >   			fmt.Printf("%d*%d=%d", i, j, i*j)
  >   		}
  >   		fmt.Println()
  >   	}
  >   }
  >   ```
  >
  >   * 因为内层循环的 j 是 <= i , 因此刚开始 j只能执行一次
  >   * 随着 i 的不断增加 , j <=i 的情况也在增加 , 因此列数增加
  >   * 外层循环控制行数
  >
  > * 绘制出一个正三角形状
  >
  >   ```go
  >   func test3() {
  >   	fmt.Println()
  >   	for i := 5; i >= 0; i-- {
  >   		for j := 4; j >= i-1; j-- {
  >   			fmt.Printf("* ")
  >   		}
  >   		fmt.Println()
  >   	}
  >   }
  >   ```

* break 和 continue

  > break会强行退出循环,会彻底退出循环,哪怕我们的循环条件是满足的
  >
  > ```go
  > for i:=0 ; true ; i++{
  >   fmt.Println(i)
  >   if i == 20 {
  >     break
  >   }
  > } 
  > ```

  > continue 会跳出本次循环,继续执行下一次循环
  >
  > ```go
  > for i:=0 ; i<=100 ; i++{
  >   if i%5 == 0{
  >     continue
  >   }
  > }
  > // 因此结果里只要是 5 的倍数,就不被打印
  > ```
  >
  > * 如果涉及到多层的循环嵌套, break和continue结束的是它所在的那层循环,不会渗透到外层
  >
  > * 但是如果想要结束外层循环,可以给外层 循环起一个名字 ,通过`break外层循环名`来结束
  >
  >   ```go
  >   name:for ... {
  >   	for ...{
  >   		if ...{
  >   		  break name
  >   		}
  >   	}
  >   }
  >   ```

###### goto语句

* goto

  > 可以无条件地将代码跳转到某处
  >
  > ```go
  > func test9(){
  >   LOOP:
  >   fmt.Printf("1")
  >   fmt.Printf("2")
  >   fmt.Printf("3")
  >   fmt.Printf("4")
  >   goto LOOP
  >   fmt.Printf("5")
  > }
  > ```
  >
  > * 当代码执行完毕打印"4"后,遇到goto,回到了LOOP所在的位置
  >
  > * 因此此段代码会循环往复地执行打印 1,2,3,4
  >
  > * goto一定慎用,并且得配合一些条件来使用,不然会出问题
  >
  >   
  >
  > goto可多用于错误的处理,一般如果捕获到错误,可以使用goto跳到错误处理