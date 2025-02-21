package main

func Switch(status int) string {
	switch status {
	case 0:
		return "初始化"
	case 1:
		return "执行中"
	case 2:
		return "重试"
	default:
		return "未知"
	}
}
func Switch1(age int) string {
	switch {
	case age >= 18:
		return "成年"
	case age < 6:
		return "幼儿"
	default:
		return "上学"
	}
}
