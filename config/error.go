package config

import "fmt"

// 自定义Json错误
type myJsonErr struct {
	err string
}

// 实现errors接口的Error方法
func (m *myJsonErr) Error() string {
	return fmt.Sprintf("%s", m.err)
}

// 生成无json数据的错误
func NoJsonDataErr() *myJsonErr {
	return &myJsonErr{err: "no jsonData"}
}
