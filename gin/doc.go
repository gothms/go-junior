package main

/*
Gin 自带的 validation
	https://gin-gonic.com/zh-cn/docs/examples/binding-and-validation/
		binding:https://github.com/go-playground/validator
		bind 依赖于反射机制
	gin 依赖 validator 库
		但是 gin 改为了 binding
	示例
		Email string `json:"email" binding:"email"`
	弊端
		如 gin 没有提供密码校验，这时最好是自己校验，好过于扩展 gin
		输出国际化的错误信息，比较受限于地区、货币、语言习惯（左到右/右到左）等

*/
