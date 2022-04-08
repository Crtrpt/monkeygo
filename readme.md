# a handy language
```
  ___              ___
 (o o)            (o o)
(  V  ) monkeygo (  V  )
--m-m--------------m-m--
```
writing an interpreter in go

[![Go](https://github.com/Crtrpt/monkeygo/actions/workflows/go.yml/badge.svg)](https://github.com/Crtrpt/monkeygo/actions/workflows/go.yml)


### 特性
-  整数 integers
-  布尔类型 bools
-  字符串 strings
-  数组 arrays
-  哈希 maps
-  前缀中缀索引操作符 prefix-, infix- and index operators
-  条件  conditionals
-  环境作用域     global and local binding
-  first-class functions
-  返回语句 return statements
-  闭包 closures

### 示例
```
let five = 5; 
let ten = 10; 

let add = fn(x, y) { 
x + y; 
}; 

let result = add(five, ten);

"foobar"

"foo bar"

[1, 2];

{"foo": "bar"}

macro(x, y) { x + y; };
```
### 特性

## 如何运行
暂不提供二进制版本 请源码运行
```
go run main.go
```
语法参考 测试用例 和源码
