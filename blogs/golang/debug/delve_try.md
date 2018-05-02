#Debug Golang With Delve

Golang在其官方[文档](https://golang.google.cn/doc/gdb) 说明

> Note that Delve is a better alternative to GDB when debugging Go programs built with the standard toolchain. It understands the Go runtime, data structures, and expressions better than GDB. Delve currently supports Linux, OSX, and Windows on amd64. For the most up-to-date list of supported platforms, please see the Delve documentation.

相比于GDB，Delve更适合于Golang程序的Debug操作，为了这句官方的推荐，我们也应该去尝试一下Delve。当前Delve可以支持Linux/OSX/Windows这几个主流平台，基本可以替代GDB的使用。

## 安装
安装Delve可以参考[官方文档](https://github.com/derekparker/delve/tree/master/Documentation/installation),这里以Mac为例。

首先确保xcode以及其插件已经安装好：

    xcode-select --install

Delve本身也是用Golang写的，所以接着通过`go get`安装Delve。

$ go get -u github.com/derekparker/delve/cmd/dlv

安装好之后，通过

    dlv version


可以看到当前版本号：

    [cz@air_11:delve]$dlv version
    Delve Debugger
    Version: 1.0.0
    Build: 8ce88095c6ea40a1d10ac2e07b7ce950f6dfaa2f

## 调试程序

现在开始写个简单求和程序

	package main

	import (
		"fmt"
	)

	func sum(a, b int) int {
		s := a + b
		return s
	}

	func main() {
		s := sum(1, 2)
		fmt.Printf("1 + 2 = %d \n", s)
	}

在代码目录下执行：

    dlv debug

看到：

    [cz@air_11:delve]$dlv debug
    Type 'help' for list of commands.
    (dlv)

进入到debug界面，这里基本和`gdb abc_exe`一样。这里会自动编译当前目录中的main包并类似gcc一样加上`-g`选项。

首先在main函数的地方设个断点：

    (dlv) b main.main
    Breakpoint 1 set at 0x10b4228 for main.main() ./main.go:18

然后，执行到该断点：

	(dlv) c
	> main.main() ./main.go:18 (hits goroutine(1):1 total:1) (PC: 0x10b4228)
	Warning: debugging optimized function
		13:	func sum(a, b int) int {
		14:		s := a + b
		15:		return s
		16:	}
		17:
	=>  18:	func main() {
		19:		s := sum(1, 2)
		20:		fmt.Printf("1 + 2 = %d \n", s)
		21:	}


这里"b"是"break"的缩写，"c"是"continue" 的缩写。如果熟悉GDB（做后台的别说不熟悉）这里会
有点不同，这里没有"run"命令，而直接是"continue"到第一个断点。如果不带参数的话，其实"run"和“continue”并没有使用上的区别。

在设置断点到sum函数，并运行到该断点：

	(dlv) b main.sum
	Breakpoint 2 set at 0x10b41d0 for main.sum() ./main.go:13
	(dlv) c
	> main.sum() ./main.go:13 (hits goroutine(1):1 total:1) (PC: 0x10b41d0)
	Warning: debugging optimized function
		 8:
		 9:	import (
		10:		"fmt"
		11:	)
		12:
	=>  13:	func sum(a, b int) int {
		14:		s := a + b
		15:		return s
		16:	}
		17:
		18:	func main() {

这里看到执行到了sum函数。

接着单步进入函数内：

	(dlv) s
	> main.sum() ./main.go:14 (PC: 0x10b41e7)
	Warning: debugging optimized function
		 9:	import (
		10:		"fmt"
		11:	)
		12:
		13:	func sum(a, b int) int {
	=>  14:		s := a + b
		15:		return s
		16:	}
		17:
		18:	func main() {
		19:		s := sum(1, 2)


在函数内部，单步下一句：

	(dlv) n
	> main.sum() ./main.go:15 (PC: 0x10b41f5)
	Warning: debugging optimized function
		10:		"fmt"
		11:	)
		12:
		13:	func sum(a, b int) int {
		14:		s := a + b
	=>  15:		return s
		16:	}
		17:
		18:	func main() {
		19:		s := sum(1, 2)
		20:		fmt.Printf("1 + 2 = %d \n", s)

如果想查看相关代码。

	(dlv) l 13
	Showing /Users/cz/Proj/golang/src/test.air/delve/main.go:13 (PC: 0x10b41d0)
	   8:
	   9:	import (
	  10:		"fmt"
	  11:	)
	  12:
	  13:	func sum(a, b int) int {
	  14:		s := a + b
	  15:		return s
	  16:	}
	  17:
	  18:	func main() {

查看指定行号，或者查看函数：

	(dlv) l sum
	Showing /Users/cz/Proj/golang/src/test.air/delve/main.go:13 (PC: 0x10b41d0)
	   8:
	   9:	import (
	  10:		"fmt"
	  11:	)
	  12:
	  13:	func sum(a, b int) int {
	  14:		s := a + b
	  15:		return s
	  16:	}
	  17:
	  18:	func main() {


最后退出：

    quit

总的来说。整个操作过程和GDB基本类似。只要熟悉Debug过程，可以很快的上手使用。

## 指定debug程序

上面的`dlv debug`是直接编译并开始debug。还可以直接传递包的路径，来指定debug那个程序。

比如：

    dlv debug github.com/nsqio/nsq/apps/nsqd

不论在哪个位置，执行上面的命令都会去debug nsqd程序。

因为Golang工具链的test和trace属性，Delve还支持debug 指定的test和trace程序，其实就是和Golang的test机制很好的结合在一起。

比如在相关目录执行：

    dlv test

就可以调试这个目录下的单元测试代码，即"_test.go"的代码。同样如果传递一个包目录，就会执行指定目录下的test. 而`dlv trace`则是执行trace代码。



## 使用vim 插件
除了在命令行里执行，Delve还提供了诸多主流编辑器的Debug插件，帮你改造你的IDE。

这里示例Vim的插件。在Delve文档中Vim插件有三个，详见[文档](https://github.com/derekparker/delve/blob/master/Documentation/EditorIntegration.md)。这里选择了[Vim-Delve](https://github.com/sebdah/vim-delve)。

因为我是用的Vundle管理Vim插件，所以在vimrc里面增加：

    Plugin 'Shougo/vimshell'
    Plugin 'Shougo/vimproc'
    Plugin 'sebdah/vim-delve'

然后执行:

    :BundleInstall

即可完成安装。

Command	| 意义
---|---
DlvAddBreakpoint |	设置断点
DlvAddTracepoint |	在这一行设置tracepoint 
DlvAttach <pid> [flags] | 	Attach 到一个运行中的程序
DlvClearAll	| 清楚所有断点
DlvCore <bin> <dump> [flags] |  调试CoreDump
DlvDebug [flags] | 对应Debug命令，Debug 程序
DlvExec <bin> [flags] | 对应exec命令
DlvRemoveBreakpoint	| 清除这一行的断点
DlvRemoveTracepoint	| 清除这一行的 tracepoint
DlvTest [flags]	 | 对应test命令
DlvToggleBreakpoint	 | 设置或者取消当行断点
DlvToggleTracepoint	 | 设置或者取消一个tracepoint 

在Vim中可以通过上面的命令随意设置断点和清除断点。然后通过"DlvDebug"或其他命令开始调试。

Vim-Delve会新建一个shell回话执行Delve命令，就和在普通的shell里面执行一样。
