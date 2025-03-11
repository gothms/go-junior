package config

/*
编译标签
	//go:build k8s
		只有当编译的时候，传了 k8s 标签，就用 k8s.go 文件，否则就不编译该文件
	使用
		go build -tags=k8s -o webook .

*/
