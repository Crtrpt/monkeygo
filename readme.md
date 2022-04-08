# a handy language
```
  ___              ___
 (o o)            (o o)
(  V  ) monkeygo (  V  )
--m-m--------------m-m--
```

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


## 需要做的
- 文档
- 包管理


# 有用的资源
- 8cc -一个很小的c语言编译器 - https://github.com/rui314/8cc
- chibicc 和8cc 相同的作者 https://github.com/rui314/chibicc
- GNU Guile 2.2 - https://www.gnu.org/software/guile/download/
- MoarVM - 一个现代化的perl vm - https://github.com/MoarVM/MoarVM
- The Lua Programming Language - lua编程语言 - https://www.lua.org/versions.html
- The Ruby Programming Language - ruby编程语言 https://github.com/ruby/ruby
- wren - wren 编程语言 https://github.com/wren-lang/wren
- c4 - C in four functions - https://github.com/rswier/c4
- tcc - Tiny C Compiler - https://github.com/LuaDist/tcc
- antlr - 解析器生成器 - https://github.com/antlr/antlr4
- Lezer - CodeMirror使用的解析器系统    https://lezer.codemirror.net/ 
- quickjs - javascript 解析器 https://github.com/bellard/quickjs
