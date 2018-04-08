# Debugging Go Code with GDB [译]

>> 一篇很老的文章，最近翻来看，好像也没人翻译，随手翻译一遍，文章来自[Golang's Blog](https://golang.google.cn/doc/gdb)

以下说明适用于Golang的标准工具链（Go编译器如gc以及其他工具），GccGo有他自己的gdb支持。

需要注意的是，对于使用标准工具链构建的Golang来说[Delve](https://github.com/derekparker/delve)相比如GDB是一个更好的选择，Delve可以更好的理解Go的运行时、数据结构以及表达式等。当前Delve可以支持Linux、OSX以及arm64平台下的Windows，最新能支持的平台列表参见[Delve的文档](https://github.com/derekparker/delve/tree/master/Documentation/installation)

GDB并不能很好理解Go程序，比如栈管理、线程以及包含了和传统GDB执行模型不一样的运行时，即便是通过gccgo来编译的程序有时候也会产生让人迷惑的信息。总的来说，虽然GDB可以在一些场景（如调试Cgo代码或者调试运行时）起到定位问题的作用，但它不是Go赖以生存的调试器，尤其是对于并发场景。或者说GDB不是Go程序首选的调试器。

所以，下面的篇幅只是当你用GDB时候的一个指引，但是并不保证一定成功。除此之外，还可以参考[GDB手册](https://sourceware.org/gdb/current/onlinedocs/gdb/)

## 简介
当在Linux/Mac OS X/FreeBSD 或者NetBSD等系统上通过gc工具链编译Go程序构建出来的二进制文件包含了 DWARFv4 调试信息可以用于GDB(需要版本大于等于7.5）调试一个运行中的进程或者Core文件。

在连接的时候，可以传递"-w"选项来省略调试信息（举例：go build -ldflags=-w prog.go）。

gc编译器生成的代码在每一行包含了函数的调用和注册表信息，这些选项有时候会使得通过gdb调试起来变得困难，所以如果需要去除这些优化的话，可以在构建时使用` go build -gcflags=all="-N -l".`

如果想通过GDB调试一个程序的core文件，需要在程序崩溃的时候触发生成一个dump文件，此时需要设置环境变量`GOTRACEBACK=crash`（更多信息参考[runtime package documentation](https://golang.google.cn/pkg/runtime/#hdr-Environment_Variables)）。

## 一般操作

* 显示文件代码或指定行号代码并设置和取消断点：


        (gdb) list
		(gdb) list line
		(gdb) list file.go:line
		(gdb) break line
		(gdb) break file.go:line
		(gdb) disas

* 显示断点和栈信息

        (gdb) bt
		(gdb) frame n

* 在栈帧中显示本地变量、参数、返回值的名称、类型位置等

        (gdb) info locals
		(gdb) info args
		(gdb) p variable
		(gdb) whatis variable

* 显示全局变量的名称、类型和位置

        (gdb) info variables regexp

## Go扩展

GDB最新的扩展机制可以让它加载指定二进制文件中的扩展脚本。工具链通过这个方法扩展了GDB来支持一些调试运行时（比如Goroutine)以及打印内建的map/slice以及channel类型的命令。

* 打印string/slice/map/channel 或接口

        (gdb) p var

* 求string/slice/map的"len"和"cap()"函数

        (gdb) p $len(var)

* 动态将接口转换为其他类型函数

        (gdb) p $dtype(var)
        (gdb) iface var

>> 如果接口的长名称不同于短名称，GDB就无法动态的找到接口值的类型。

* 查看goroutines:

        (gdb) info goroutines
		(gdb) goroutine n cmd
		(gdb) help goroutine


例子：

        (gdb) goroutine 12 bt

如果想探寻其工作流程，或者想扩展的话，可以查看Go源码目录中的"src/runtime/runtime-gdb.py"。这里的脚本依赖了一些连接器（src/cmd/link/internal/ld/dwarf.go）在DWARF中保留的 特殊的魔数类型(hash<T,U>) 和变量。

如果想了解debug信息长什么样。可以通过运行`objdump -W a.out` 并浏览 ".debug_*"段。

### 关键点

1. 这里的完美打印只能打印string，但是继承与string的类型不行
2. Runtime里面C的类型信息是没有的
3. GDB无法完全理解Go的名称，并将"fmt.Print"理解成需要用双引号括起来的非结构化字符串，对于类似"pkg.(*MyType).Meth"这种的就更难理解了。
4. 所有的全局对象都在"main"包里面

## 教程
在这个教程中，我们调试"[regexp](https://golang.google.cn/pkg/regexp/)"包中的测试程序。首先去到"$GOROOT/src/regexp"目录下然后运行"go test -c"，然后生成可执行文件"regexp.test"。

### 开始

运行GDB并调试regexp.test:

    $ gdb regexp.test
    GNU gdb (GDB) 7.2-gg8
	Copyright (C) 2010 Free Software Foundation, Inc.
	License GPLv  3+: GNU GPL version 3 or later <http://gnu.org/licenses/gpl.html>
	Type "show copying" and "show warranty" for licensing/warranty details.
	This GDB was configured as "x86_64-linux".

    Reading symbols from  /home/user/go/src/regexp/regexp.test...
	done.
	Loading Go Runtime support.
	(gdb)

"Loading Go Runtime support"表示GDB从"$GOROOT/src/runtime/runtime-gdb.py."加载了扩展脚本。

通过传递"-d"和"$GOROOT",可以使GDB找到Go运行时以及对应的脚本文件。

    $ gdb regexp.test -d $GOROOT

如果因为什么原因，GDB始终找不到这个目录和脚本，那么可以在GDB里面来加载。（假设你的go源码代码在"~/go/"）

    (gdb) source ~/go/src/runtime/runtime-gdb.py
	Loading Go Runtime support.


### 查看源文件

使用"l"或者"list"命令查看代码

    (gdb) l

通过给"list"传递一个函数名（必须包含包名），可以查看指定位置的源代码。

    (gdb) l main.main

查看指定文件的某行

    (gdb) l regexp.go:1
    (gdb) # Hit enter to repeat last command. Here, this lists next 10 lines.

### 命名
函数名和变量名必须要加上其所在包的包名。regexp包中的额函数"Compile"在GDB中需要写成"regexp.Compile"。

方法名必须要加上他的类的类型。比如"* Regexp"类型的 "String" 方法要写成 "regexp.(*Regexp).String"。

被其他变量隐藏的变量需要在其前面增加一个数字索引，而被闭包🎵的变量则需要用指针操作符"&"来前缀。

设置断点，在"TestFind"函数处设置一个断点：

    (gdb) b 'regexp.TestFind'
	Breakpoint 1 at 0x424908: file /home/user/go/src/regexp/find_test.go, line 148.
	Run the program:

    (gdb) run
	Starting program: /home/user/go/src/regexp/regexp.test

    Breakpoint 1, regexp.TestFind (t=0xf8404a89c0) at /home/user/go/src/regexp/find_test.go:148
	148	func TestFind(t *testing.T) {

此时停在断点出，可以查看是哪个goroutines在运行以及在做什么：


    (gdb) info goroutines
	1  waiting runtime.gosched
	* 13  running runtime.goexit

"*" 标记表示当前在的goroutine。

### 查看堆栈
查看暂停时的堆栈信息：

    (gdb) bt  # backtrace
	#0  regexp.TestFind (t=0xf8404a89c0) at /home/user/go/src/regexp/find_test.go:148
	#1  0x000000000042f60b in testing.tRunner (t=0xf8404a89c0, test=0x573720) at /home/user/go/src/testing/testing.go:156
	#2  0x000000000040df64 in runtime.initdone () at /home/user/go/src/runtime/proc.c:242
	#3  0x000000f8404a89c0 in ?? ()
	#4  0x0000000000573720 in ?? ()
	#5  0x0000000000000000 in ?? ()

另一个序号为1的线程停在"runtime.gosched"线程中，阻塞了一个chanel：

    (gdb) goroutine 1 bt
	#0  0x000000000040facb in runtime.gosched () at /home/user/go/src/runtime/proc.c:873
	#1  0x00000000004031c9 in runtime.chanrecv (c=void, ep=void, selected=void, received=void)
	at  /home/user/go/src/runtime/chan.c:342
	#2  0x0000000000403299 in runtime.chanrecv1 (t=void, c=void) at/home/user/go/src/runtime/chan.c:423
	#3  0x000000000043075b in testing.RunTests (matchString={void (struct string, struct string, bool *, error *)}
	0x7ffff7f9ef60, tests=  []testing.InternalTest = {...}) at /home/user/go/src/testing/testing.go:201
	#4  0x00000000004302b1 in testing.Main (matchString={void (struct string, struct string, bool *, error *)} 
	0x7ffff7f9ef80, tests= []testing.InternalTest = {...}, benchmarks= []testing.InternalBenchmark = {...})
	at /home/user/go/src/testing/testing.go:168
	#5  0x0000000000400dc1 in main.main () at /home/user/go/src/regexp/_testmain.go:98
	#6  0x00000000004022e7 in runtime.mainstart () at /home/user/go/src/runtime/amd64/asm.s:78
	#7  0x000000000040ea6f in runtime.initdone () at /home/user/go/src/runtime/proc.c:243
	#8  0x0000000000000000 in ?? ()

栈帧显示我们正在执行"regexp.TestFind"函数：

    (gdb) info frame
	Stack level 0, frame at 0x7ffff7f9ff88:
	 rip = 0x425530 in regexp.TestFind (/home/user/go/src/regexp/find_test.go:148); 
         saved rip 0x430233
	called by frame at 0x7ffff7f9ffa8
	source language minimal.
	Arglist at 0x7ffff7f9ff78, args: t=0xf840688b60
	Locals at 0x7ffff7f9ff78, Previous frame's sp is 0x7ffff7f9ff88
	Saved registers:
	 rip at 0x7ffff7f9ff80


命令`info locals`列出了这个函数的所有的局部变量和值，因为它也会打印未经初始化的变量，所以使用时会有一定的危险，因为未经初始化的slice会导致gdb去打印任意长度的数组。

函数参数：

    (gdb) info args
	t = 0xf840688b60

注意，这里打印参数时是打印的一个Regexp的指针。GDB在类型名的右边放量一个"*"来修饰一个"struct"，跟传统的C风格一样。

    (gdb) p re
    (gdb) p t
    $1 = (struct testing.T *) 0xf840688b60
	(gdb) p t
	$1 = (struct testing.T *) 0xf840688b60
	(gdb) p *t
	$2 = {errors = "", failed = false, ch = 0xf8406f5690}
	(gdb) p *t->ch
	$3 = struct hchan<*testing.T>

"struct hchan<*testing.T>"是运行时内部的channel的表示，当前是空的，否则GDB会打印他的内容。

接着往下走：

    (gdb) n  # execute next line
	149             for _, test := range findTests {
	(gdb)    # enter is repeat
	150                     re := MustCompile(test.pat)
	(gdb) p test.pat
	$4 = ""
	(gdb) p re
	$5 = (struct regexp.Regexp *) 0xf84068d070
	(gdb) p *re
	$6 = {expr = "", prog = 0xf840688b80, prefix = "", prefixBytes =  []uint8, prefixComplete = true, 
	    prefixRune = 0, cond = 0 '\000', numSubexp = 0, longest = false, mu = {state = 0, sema = 0}, 
        machine =  []*regexp.machine}
    (gdb) p *re->prog
	$7 = {Inst =  []regexp/syntax.Inst = {{Op = 5 '\005', Out = 0, Arg = 0, Rune =  []int}, {Op = 
        6 '\006', Out = 2, Arg = 0, Rune =  []int}, {Op = 4 '\004', Out = 0, Arg = 0, Rune =  []int}}, 
        Start = 1, NumCap = 2}


通过"s"可以进入导函数"String"

    (gdb) s
	regexp.(*Regexp).String (re=0xf84068d070, noname=void) at /home/user/go/src/regexp/regexp.go:97
	97      func (re *Regexp) String() string {

看下所在位置的堆栈

    (gdb) bt
	#0  regexp.(*Regexp).String (re=0xf84068d070, noname=void)
        at /home/user/go/src/regexp/regexp.go:97
	#1  0x0000000000425615 in regexp.TestFind (t=0xf840688b60)
        at /home/user/go/src/regexp/find_test.go:151
	#2  0x0000000000430233 in testing.tRunner (t=0xf840688b60, test=0x5747b8)
        at /home/user/go/src/testing/testing.go:156
	#3  0x000000000040ea6f in runtime.initdone () at /home/user/go/src/runtime/proc.c:243
	....

看代码：

    (gdb) l
	92              mu      sync.Mutex
	93              machine []*machine
	94      }
	95
	96      // String returns the source text used to compile the regular expression.
	97      func (re *Regexp) String() string {
	98              return re.expr
	99      }
	100
    101     // Compile parses a regular expression and returns, if successful,


### 完美打印

GDB的完美打印机制通过 regexp匹配类型名，比如slices:

    (gdb) p utf
	$22 =  []uint8 = {0 '\000', 0 '\000', 0 '\000', 0 '\000'}

slices/array以及strings的切片不是C指针。GDB不能为你操作子脚本，但是你可以通过查看运行时内部的名称表示来查看（通过TAB自动补全）


    (gdb) p slc
	$11 =  []int = {0, 0}
	(gdb) p slc-><TAB>
	array  slc    len    
	(gdb) p slc->array
	$12 = (int *) 0xf84057af00
	(gdb) p slc->array[1]
	$13 = 0


扩展函数"$len"和"$cap"可以作用在strings/arrays和sclices:

    (gdb) p $len(utf)
	$23 = 4
	(gdb) p $cap(utf)
	$24 = 4

Channels和Maps是引用类型，GDB会按照C++格式的指针进行打印如“hash<int,string>*”。解引用也可以完美支持。

引用在运行时里面被表示为一个执行类型的指针和一个执行值的指针。Go的GDB扩展会自动将这两者解码并在进行打印。扩展函数"$dtype"将解码动态类型（比如例子中的regexp.go的293行）

    (gdb) p i
	$4 = {str = "cbb"}
	(gdb) whatis i
	type = regexp.input
	(gdb) p $dtype(i)
	$26 = (struct regexp.inputBytes *) 0xf8400b4930
	(gdb) iface i
	regexp.input: struct regexp.inputBytes *

