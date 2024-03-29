---
title: "Awesome Rust"
---
# 工程结构
## 格式化

### rustfmt.toml
用于制定rust格式化的规则

# 类库
## Console
### [clap](https://github.com/clap-rs/clap)
命令行解析框架，做Flag解析

### [console](https://github.com/mitsuhiko/console)
多终端统一样式输出,支持Unicod和ANSI,有颜色、字体，甚至连Emoj表情都有

### [indicatif](https://github.com/mitsuhiko/indicatif)
终端展示进度条，拥有各种风格

### [tui](https://github.com/fdehau/tui-rs)
console的富UI终端显示

## Net
### [net2](https://docs.rs/net2/latest/net2/)
标准库net的补充，用于常见的tcp/udp socket操作

### [socket2](https://crates.io/crates/socket2)
net已经废弃了，新的用socket2

## HTTP
### [reqwest](https://docs.rs/reqwest)
HTTP Client的一种实现

### [hyper](https://crates.io/crates/hyper)
更快的Http实现，包含了Http1和Http2

## Encoding

### [Bytes](https://docs.rs/bytes)
byte操作

### [Serde JSON](https://github.com/serde-rs/json)
Rust的JSON库

### [bincode](https://docs.rs/bincode)
类似Go里面的gob，一种二进制序列化策略


## Crypto
### [ed25519-dalek](https://docs.rs/ed25519-dalek)
ed25519 生成key、校验等功能

## Random
### [rand](https://github.com/rust-random/rand)
非标准库中的随机数

## FileSystem
### [dirs](https://github.com/dirs-dev/dirs-rs)
获取系统的home目录、audio目录等

## WebFramework
### [actix](https://github.com/actix/actix)
异步的web框架，基于tokio

### [hyper](https://docs.rs/hyper)
Rust实现的HTTP服务器，类似go的fasthttp

## DB
### [rocksdb](https://github.com/rust-rocksdb/rust-rocksdb)
区块链都在用的rocksdb的rust的binding，是一个binding，不是rust重新实现


## Trace
### [tracing](https://docs.rs/tracing)
tokio出品，trace工具

## Syntax
### [syn](https://docs.rs/syn)
解析Rust的Token到Rust的语法书，比如在宏中使用

## BlockChain

### [tiny-bip39](https://docs.rs/tiny-bip39/0.8.2/bip39/)
BIP39的Rust实现

### [bs58](https://docs.rs/bs5)
Base58实现

## Log
### [env_logger](https://docs.rs/env_logger)
通过环境变量控制的日志

## Cache
### [cached](https://docs.rs/cached)
buffer 缓存

## Utils

### [git-version](https://github.com/fusion-engineering/rust-git-version)
获取仓库当前的git commit等相关信息作为版本号

### [rustc-version](https://github.com/djc/rustc-version-rs)
获得当前rustc的版本

### [num_cpus](https://github.com/seanmonstar/num_cpus)
获得当前机器上有多少个CPU核

### [self_update](https://docs.rs/self_update/latest/self_update/)
工程子更新，根据github信息更新到最新tag





