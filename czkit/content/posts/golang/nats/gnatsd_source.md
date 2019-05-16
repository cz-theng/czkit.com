---
title: "NATS 开源学习"
date: 2019-03-02T22:07:36+08:00
categories:
  - "golang"
tags:
  - "NATS"
  - "MQ"

description: "NATS源码学习系列文章基于[gnatsd1.0.0](https://github.com/nats-io/gnatsd/tree/v1.0.0)"
---

# NATS 开源学习——0X00：协议

> NATS源码学习系列文章基于[gnatsd1.0.0](https://github.com/nats-io/gnatsd/tree/v1.0.0)。该版本于2017年7月13
> 日发布（[Release v1.0.0](https://github.com/nats-io/gnatsd/releases/tag/v1.0.0)）,在此之前v0.9.6是2016年12月
> 16日发布的,中间隔了半年。算是一个比较完备的版本，但是这个版本还没有增加集群支持。为什么选择这个版本呢？
> 因为一来这个版本比较稳定，同时也包含了集群管理和[Stream](https://github.com/nats-io/nats-streaming-server)
> 落地相关的逻辑，相对完善。

## gnatsd
在好多年前写过一篇关于NATS的初体验[NATS之gnatcd初体验](http://www.czkit.com/posts/golang/nats/have_a_try_with_gnatcd/)。
那个时候gnatsd才0.8.0。当时只是因为新看到一个听说性能很屌的MQ，于是就尝试了一下，一晃眼，gnatsd都发布了n多个版本了（1.4.x都出来了）
也成为了[CNCF](https://www.cncf.io/blog/2018/03/15/cncf-to-host-nats/)孵化项目之一。所以通过gnatsd的一个稳定版本来学习下
这个MQ的实现。

<!--more-->


NATS定义了一套非常简单的协议，来实现一个基于TCP链接的发布订阅系统。其核心就是订阅与发布，为此使用Golang实现了这套协议转发功能的gnatsd服务。这里
是说得益于Golang呢？还是说项目方写程序能力厉害呢？还是服务器设计的厉害呢？反正项目方原先是有一个Ruby的实现，现在抛弃了，现在官方主要维护的是这个Golang版本的gnatsd。

把gnatsd的源码down下来后，其核心代码只有不到一万行，另外一个MQ项目[NSQ](https://github.com/nsqio/nsq)也是类似的规模。 这些高效的Borkerless MQ是如何通过Golang来实现？通过看代码，这里进行了些分析。

![](../images/gitbook.jpg)

目录如上，主要通过协议分析，再到协议的实现去看gnatsd是如何实现PUB/SUB的功能。

gnats与2017年7月13日发布[Release v1.0.0](https://github.com/nats-io/gnatsd/releases/tag/v1.0.0)
，而此之前v0.9.6是2016年12月16日发布的,中间隔了半年。算是一个比较完备的版本。

故此通过对1.0版本代码的学习，写了篇总结，详见[NATS开源学习](https://cz-it.gitbook.io/nats-source/)


