## Table of Contents
- [Background](#background)
- [Install](#install)
- [Usage](#usage)
- [Features](#features)


### Background

### Install




### Usage

```go

import "github.com/Celephaiss/zapWrapper"


// 只需要Init一次
zapWrapper.Init("./test.log", "debug")

// 在不同的goroutine里面调用NewSugar(name)获得同一个logger，
// 可以通过传入不同的name来标识goroutine
// 通过logger写日志是并发安全的。
l1 := zapWrapper.NewSugar("hello")
l1.Error("this is a test")

// 2021-05-12T15:47:26.904+0800	error	hello	zapLogger/logger_test.go:23	this is a test


// l1和l2都是往test.log里面写日志。
l2 := zapWrapper.NewSugar("hello2")
l2.Error("this is another test")

// 2021-05-12T15:47:26.904+0800	error	hello2	zapLogger/logger_test.go:23	this is another test


```


### Features

1. 日志切割
2. 将不同级别的日志输出到不同文件

