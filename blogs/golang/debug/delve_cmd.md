# Delve Commands
前面的[Debug Golang With Delve](http://www.czkit.com/posts/golang/debug/delve_try/)介绍了Delve的基本用法。
其实Delve提供了和gdb基本一致的操作指令。除了可以像gdb一样“尬调”,Delve还可以通过
Vim来触发Delve 命令加速调试。


## Delve基本指令

### break : 设置断点

"break" 就和gdb的基本一致，可以对行号设断点，也可以对方法设断点。可以敲其缩写"b"。

### breakpoints 打印所有断点

"breakpoints" 可以打印出所有激活的断点，类似Xcode里面断点界面中把所有的断点列出来一样。可以缩写为"bp",如：

        (dlv) bp
        Breakpoint unrecovered-panic at 0x102da10 for runtime.startpanic() /Users/apollo/go/src/runtime/panic.go:577 (0)
            print runtime.curg._panic.arg
        Breakpoint 1 at 0x113a418 for github.com/golang/groupcache/lru.(*Cache).Add() ./lru.go:56 (1)
        Breakpoint 2 at 0x113a9b3 for github.com/golang/groupcache/lru.(*Cache).Remove() ./lru.go:86 (0)
        Breakpoint 3 at 0x113aa9f for github.com/golang/groupcache/lru.(*Cache).RemoveOldest() ./lru.go:96 (0)
        Breakpoint 4 at 0x113a833 for github.com/golang/groupcache/lru.(*Cache).Get() ./lru.go:74 (0)

这里总共有四个断点。

### condition: 条件断点

    condition <breakpoint name or id> <boolean expression>

使一个断点仅在后面的布尔表达式为真实才真正断下来。比如我们希望上面的Add函数只在key不为nil的时候才断下来：

    (dlv) condition 1 key != nil

### continue : 继续执行到下一个断点
continue和gdb的没什么区别，缩写为"c"


### list : 显示代码
走到断点的地方，可以直接"list"命令显示上下代码行，缩写为"l"。

也可以给list传递行号。比如：

        (dlv) l lru.go:50
        Showing /Users/apollo/go_proj/src/github.com/golang/groupcache/lru/lru.go:50 (PC: 0x113a27a)
          45:	// If maxEntries is zero, the cache has no limit and it's assumed
          46:	// that eviction is done by the caller.
          47:	func New(maxEntries int) *Cache {
          48:		return &Cache{
          49:			MaxEntries: maxEntries,
          50:			ll:         list.New(),
          51:			cache:      make(map[interface{}]*list.Element),
          52:		}
          53:	}
          54:
          55:	// Add adds a value to the cache.

### args：查看函数参数

        (dlv) b github.com/golang/groupcache/lru.(*Cache).Add
        Breakpoint 1 set at 0x113a418 for github.com/golang/groupcache/lru.(*Cache).Add() ./lru.go:56
        (dlv) c
        > github.com/golang/groupcache/lru.(*Cache).Add() ./lru.go:56 (hits goroutine(5):1 total:1) (PC: 0x113a418)
        Warning: debugging optimized function
            51:			cache:      make(map[interface{}]*list.Element),
            52:		}
            53:	}
            54:
            55:	// Add adds a value to the cache.
        =>  56:	func (c *Cache) Add(key Key, value interface{}) {
            57:		if c.cache == nil {
            58:			c.cache = make(map[interface{}]*list.Element)
            59:			c.ll = list.New()
            60:		}
            61:		if ee, ok := c.cache[key]; ok {
        (dlv) args
        c = (*github.com/golang/groupcache/lru.Cache)(0xc42000a080)
        key = github.com/golang/groupcache/lru.Key(string) "myKey"
        value = interface {}(int) 1234

如上面的例子，首先设置断点，然后运行到该断点后，执行"args"命令，可以看到这里的两个参数：

    key = "myKey"
    value = 1234



### check : 设置check断点
什么是"check"断点，如果用过gdb的"checkpoint"就比较好理解，其实就是可以返回到之前的执行上下文的快照，其操作和
"break"基本一致

### checkpoints :打印所有的checkpoint

和"breakpoints"类似，打印所有设置过的checkpoint


### clear: 清除普通断点

传递断点名或者ID，来清除指定通过"break"设置的断点。

### clear-checkpoint： 清除 check断点

传递断点名或者ID，来清除指定通过"check"设置的断点。

### clearall :清除所有断点

清除所有设置过的断点。如果传递了行号信息，那么仅清除该行上的断点。

### exit : 退出调试
关闭Delve，缩写为"q"

### down: 移植下一帧
在调试


## 使用vim 插件

除了在命令行里执行，Delve还提供了诸多主流编辑器的Debug插件，帮你改造你的IDE。

这里示例Vim的插件。在Delve文档中Vim插件有三个，详见[文档](https://github.com/derekparker/delve/blob/master/Documentation/EditorIntegration.md) 。 看过我前面的[Golang With Vim](http://www.czkit.com/posts/vim/vim_golang/) 
肯定知道这里我们会选择 [vim-go](https://github.com/fatih/vim-go) 。


这里我们以Golang官方的[GroupCache](https://github.com/golang/groupcache) 中的[LRU](https://github.com/golang/groupcache/tree/master/lru) 算法包为例。先来看牛逼的效果图：

![](../images/vim_delve.png)


"vim-go" 已经内置对 Delve支持。只要在编辑Go代码的是执行 `:GoDebugStart` 就可以开始调试了。

当然，这里集成到Vim中的优势，就是可以像Xcode/VS一样，一边浏览代码，一边下断点。而不用像传统后台用Gdb调C一样，
从记忆里面寻找代码行和函数名。

通过执行:`:GoDebugBreakPoint`,可以在光标所在的位置行设置断点。比如图中的位置。

这里因为我们是在一个子包中，通过单元测试来调试代码，所以设置好断点后，通过`:GoDebugTest` 来调试测试代码，就和
dlv命令行传递test命令一样。

当开始后，通过`:GoDebugContinue` 触发dlv的"continue"命令，执行到下一个断点。

到达断点后，可以通过`:GoDebugNext`触发dlv的"next"命令，往下执行一行。或者`:GoDebugStep`触发dlv的"step"命令进入
到函数中(对应的通过`:GoDebugStepOut`跳出来)。

进入到函数中可以通过`:GoDebugPrint` 触发dlv的"print"命令，打印相关变量。

最后`:GoDebugStop`可以退出调试界面。

按照dlv的操作，基本都有其对应的命令。这里可以在.vimrc中对常用的如GoDebugContinue、GoDebugStep配置快捷键。



## 附录:Delve 命令


命令| 描述 
--------|------------
[args](#args) | 打印函数参数
[break](#break) | 设置断点
[breakpoints](#breakpoints) | 打印断点信息
[check](#check) | 创建checkpoint 断点
[checkpoints](#checkpoints) | 打印断点信息
[clear](#clear) |  删除断点
[clear-checkpoint](#clear-checkpoint) | 删除 checkpoint.
[clearall](#clearall) | 删除所有断点
[condition](#condition) | 设置条件断点
[config](#config) | 修改配置参数
[continue](#continue) | 运行直到下一个断点
[disassemble](#disassemble) | 反汇编
[down](#down) | 将当前帧往下移一个
[exit](#exit) | 退出
[frame](#frame) | 设置当前帧
[funcs](#funcs) | 打印函数列表
[goroutine](#goroutine) | 显示或者切换goroutine
[goroutines](#goroutines) |罗列所有goroutine 
[help](#help) | 打印帮助信息
[list](#list) | 显示代码
[locals](#locals) | 打印局部变量
[next](#next) | 执行到下一行
[on](#on) | 触发断点时执行脚本
[print](#print) | 打印表达式结果
[regs](#regs) | 打印寄存器内容
[restart](#restart) | 重新运行
[rewind](#rewind) | 后台运行直到断点或者程序结束
[set](#set) | 修改变量内容
[source](#source) | 执行一个包含delve命令的脚本
[sources](#sources) | 罗列源码文件
[stack](#stack) | 打印栈内容
[step](#step) | 单步进函数
[step-instruction](#step-instruction) | 单步进CPU指令级别
[stepout](#stepout) | 单步出来
[thread](#thread) | 切换到指定线程
[threads](#threads) | 打印每个线程的trace信息
[trace](#trace) | 设置tracepoint.
[types](#types) | 罗列所有类型
[up](#up) | 将当前帧往上移
[vars](#vars) | 打印包内变量
[whatis](#whatis) | 打印表达式结果的类型

