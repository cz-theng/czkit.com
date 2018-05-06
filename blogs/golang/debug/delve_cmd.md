# Delve指令详解


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

