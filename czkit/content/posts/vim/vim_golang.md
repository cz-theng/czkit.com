---
title: "Golang With Vim"
date: 2019-04-27T22:07:36+08:00
categories:
  - "vim"
tags:
  - "vim"

description: "发现Vim的Golang环境配置已经被大神们整的这么Easy，如今Vundle可以让小白用户快速的配置好各种插件。所以只要会Vim基本使用"
---


在很久以前的官网[golang.org](https://golang.org/)上是有一段视频的。视频里面正是用了Vim作为编辑器。
从Golang出生时，就有人为Vim好Emacs写了扩展插件。当然现在的Sublime/VSCode/Atom都可以很方便的配置成
一款顺手的Golang开发工具。

因为经历的原因，我是从Vim到Emacs现在又回到Vim。发现Vim的Golang环境配置已经被大神们整的这么Easy，最开始用
Vim那会，是要自己去[vim.org](http://vim.org)上面找插件，然后自己放到.vim目录下，在写vimrc，稍稍不注意
就会弄的vim起来一堆错误。而今，有了[Vundle](https://github.com/VundleVim/Vundle.vim) 和[Pathogen](https://github.com/tpope/vim-pathogen)这样的插件管理神器，可以让小白用户快速的配置好各种插件。所以只要会Vim基本使用
，就可以很快的使用Vim来写Golang。

自从2006年发布了Vim7以来，直到2016年，时隔10年，Vim终于发布了Vim8，新的Vim汲取了NeoVim一些新的特性，总的
来说建议过度到Vim8上，另外目前MacOS默认的Vim7.3，所以需要自己装一下，可以使用二进制编译安装或者直接使用[MacVim](https://github.com/macvim-dev/macvim)

下面假定是重新安装的Vim，没有经过其他配置。Vim的用户配置文件在`~/.vimrc` ,一些插件和其他附加文件在`~/.vim`
目录。这里假定这个目录和文件都是不存在的。


<!--more-->


## 1. 安装Vundle
这里选择了Vundle这个Vim插件神器，下面先安装下。

首先下载Vundle到.vim目录

    git clone https://github.com/VundleVim/Vundle.vim.git ~/.vim/bundle/Vundle.vim

git会自动创建相关目录。这样我们就有了我们的Vim的配置文件目录`~/.vim`目录。

然后我们再创建配置文件`~/.vimrc` 直接用Vim创建

    vim ~/.vimrc

然后写入：

        set nocompatible              " be iMproved, required
        filetype off                  " required

        " set the runtime path to include Vundle and initialize
        set rtp+=~/.vim/bundle/Vundle.vim
        call vundle#begin()
        " alternatively, pass a path where Vundle should install plugins
        "call vundle#begin('~/some/path/here')

        " let Vundle manage Vundle, required
        Plugin 'VundleVim/Vundle.vim'


        " All of your Plugins must be added before the following line
        call vundle#end()            " required
        filetype plugin indent on    " required

所有的插件需要卸载`call vundle#end()`的前面。

然后执行:

    :PluginInstall

就可以安装相关的插件了。这里先安装好"Vundle"。


    " Installing plugins to /Users/apollo/.vim/bundle   
    . Plugin 'VundleVim/Vundle.vim'
    ...
    Done!

提示表示安装完成。

## 2. Golang插件

Golang在出生的时候就提供了Vim插件支持，并且越发展越完善，目前集成了godoc/golint/godef等等工具。同时还有个
为Golang提供自动补全功能的项目[Gocode](https://github.com/nsf/gocode)。所以先配置这两个插件：

          5 "*********BEGIN OF VUNDLE ***********************************
          6 set nocompatible              " be iMproved, required
          7 filetype off                  " required
          8
          9 " set the runtime path to include Vundle and initialize
         10 set rtp+=~/.vim/bundle/Vundle.vim
         11 call vundle#begin()
         12 " alternatively, pass a path where Vundle should install plugins
         13 "call vundle#begin('~/some/path/here')
         14
         15 " let Vundle manage Vundle, required
         16 Plugin 'VundleVim/Vundle.vim'
         17
         18 "Plugins
         19 Plugin 'nsf/gocode', {'rtp': 'vim/'} " golang
         20 Plugin 'fatih/vim-go'                " golang misc/vim
         32
         33 " All of your Plugins must be added before the following line
         34 call vundle#end()            " required
         35 filetype plugin indent on    " required
         36 "********************END OF VUNDLE **************************

然后执行:

    :PluginInstall

## 3. Tags插件

IDE一般都提供了Tags界面，可以比较方便的找到函数或者类的定义。Vim可以借助ctags和gotags实现同样的功能。

首先安装ctags,在Mac上可以直接通过Brew 来安装：

    brew install ctags

> 需要注意的是，如果你之前是个Emacs党，那么一般会安装好Emacs的ctags，而gotags不能识别，所以需要安装完ctags后
> 重新连接下。

执行 ：

    ctags --version
    Exuberant Ctags 5.8, Copyright (C) 1996-2009 Darren Hiebert
    Compiled: May  3 2018, 20:26:45
    Addresses: <dhiebert@users.sourceforge.net>, http://ctags.sourceforge.net
    Optional compiled features: +wildcards, +regex

确定安装完成。

然后在安装[GoTags](https://github.com/jstemmer/gotags):

    go get -u github.com/jstemmer/gotags

最后配置".vimrc"


         82 "**** go tags
         83 let g:tagbar_type_go = {
         84     \ 'ctagstype' : 'go',
         85     \ 'kinds'     : [
         86         \ 'p:package',
         87         \ 'i:imports:1',
         88         \ 'c:constants',
         89         \ 'v:variables',
         90         \ 't:types',
         91         \ 'n:interfaces',
         92         \ 'w:fields',
         93         \ 'e:embedded',
         94         \ 'm:methods',
         95         \ 'r:constructor',
         96         \ 'f:functions'
         97     \ ],
         98     \ 'sro' : '.',
         99     \ 'kind2scope' : {
        100         \ 't' : 'ctype',
        101         \ 'n' : 'ntype'
        102     \ },
        103     \ 'scope2kind' : {
        104         \ 'ctype' : 't',
        105         \ 'ntype' : 'n'
        106     \ },
        107     \ 'ctagsbin'  : 'gotags',
        108     \ 'ctagsargs' : '-sort -silent'
        109 \ }    

并在上面的插件中增加一个Tag现实的插件：

    Plugin 'majutsushi/tagbar'           " tagbar should install ctags&gotags

并重新安装，此时在打开go的文件后，执行`:TagbarToggle` 就可以看到如下的Tag界面了：

![tags](../images/tags.png)

## 4. 目录插件
vim 说到目录插件，有一堆的选择，最牛逼的还是[NERDTree](https://github.com/scrooloose/nerdtree).
现在有了Vundle比以前方便多了，在也不用去Vim.org了。

直接在上面的插件位置增加：

    Plugin 'scrooloose/nerdtree'

然后执行安装，执行`:NERDTreeToggle`
就可以调出目录界面：

![dir](../images/dir.png)

## 5. 自动补齐
虽然上面的gocode已经基本能满足补齐需要，不过还有个牛X的补齐插件，可以扩展，当然其Golang部分也是用的Gocode。

神器名曰: YouCompleteMe


名字都霸气侧漏。和上面一样，在.vimrc中加上插件：

    Plugin 'Valloric/YouCompleteMe'      " for auto complete

然后安装插件。

这个牛X插件是前Google大厂员工写的，其独特的地方是不是用vim脚本写的，而是用C++来实现的，所以其补齐效率非常高。

所以其安装方法也不一样。

在上面安装好了之后，还要去到

    cd ~/.vim/bundle/YouCompleteMe

执行：

    ./install.py 
    ./install.py  --go-completer

执行编译安装。

具体效果可以参照[YouCompleteMe](https://github.com/Valloric/YouCompleteMe)的动图。



## 未完待续

我自己的.vimrc可以参考[Github](https://github.com/cz-it/.vim.d)
