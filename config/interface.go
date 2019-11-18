package config

// 定义配置解析接口
type ConfParser interface {
	getConfig(k string) (v interface{}, err error)
}
