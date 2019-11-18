package api

// 定义apollo配置中心无缓存的HTTP接口
type NocachedHttper interface {
	GET(config_server_url, appId, clusterName, namespaceName, releaseKey, ip string) (j interface{}, err error)
}
