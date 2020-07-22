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

    

* 常用的fmt.Printftv占位符

  > * 查看类型`fmt.Printf("%T\n",a)`
  >
  > * 查看值`fmt.Printf("%v\n",a)`
  >
  > * 转为二进制 八进制 十进制 十六进制
  >
  >   ```go
  >   fmt.Print("%b\n",a)
  >   fmt.Print("%o\n",a)
  >   fmt.Print("%d\n",a)
  >   fmt.Print("%x\n",a)
  >   ```
  >
  > * 字符串`fmt.Printf("%s\n",a)`

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
  >     ```
  >
  > * 字符串的修改

