package config

import "strings"

// 新建应用配置类函数
func newAppConfig(e string, appid string, cluster string, ns string, key string) *appConfig {
	// apollo配置中，集群名，默认“default”
	if cluster == "" {
		cluster = apolloDefaultCluster
	}
	// apollo配置中，名称空间，默认“application”
	if ns == "" {
		ns = apolloDefaultNamespace
	}
	// 查询目标配置的键名，默认是dubbo服务注册端口：“dubbo.provider.port”
	if key == "" {
		key = defaultDubboPortKey
	}
	// 初始化应用配置类实例
	var a = &appConfig{env: e, cluster: cluster, namespace: ns, appID: appid}
	// 根据环境的不同，指定该环境对应的apollo配置中心地址
	a.apolloSocket = func() string {
		switch strings.ToLower(a.env) {
		case "uat":
			return uatApolloConfigSocket
		case "sit":
			return sitApolloConfigSocket
		case "dev":
			return devApolloConfigSocket
		case "sandbox":
			return sandboxApolloConfigSocket
		case "perform":
			return performApolloConfigSocket
		default:
			return ""
		}
	}()

	return a
}

// 获取Dubbo服务指定注册端口的函数
func GetDubboPort(e string, appid string, cluster string, ns string, key string) (port string, err error) {
	// 实例化应用配置
	appConf := newAppConfig(e, appid, cluster, ns, key)
	// 调用配置解析接口，并解析配置
	var confParser ConfParser = appConf
	conf, confErr := confParser.getConfig(apolloDefaultConfKey)
	if confErr != nil {
		err = confErr
		port = ""
		return
	}
	if conf != nil {
		portJ := conf.(map[string]interface{})[defaultDubboPortKey]
		port = portJ.(string)
	} else {
		port = ""
		err = NoJsonDataErr()
	}
	return
}
