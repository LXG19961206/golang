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

* 

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

    