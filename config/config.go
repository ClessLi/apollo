package config

import (
	"fmt"
	"github.com/ClessLi/apollo/api"
	"github.com/PuerkitoBio/goquery"
	"github.com/thedevsaddam/gojsonq"
	"io/ioutil"
	"net/http"
	"strings"
)

// 定义应用配置类
type appConfig struct {
	// 环境，apollo配置中心套接字，目标配置键名，应用id，apollo集群名，apollo命名空间，上一次的releaseKey，应用部署的ip
	env, apolloSocket, confKey, appID, cluster, namespace, releaseKey, ip string
}

// 实现应用配置解析接口的getConfig方法
func (a *appConfig) getConfig(k string) (v interface{}, err error) {
	var apolloApi api.NocachedHttper = a
	jsonData, jsonErr := apolloApi.GET(a.apolloSocket, a.appID, a.cluster, a.namespace, a.releaseKey, a.ip)
	if jsonErr != nil {
		err = jsonErr
		return
	}
	v = gojsonq.New().FromString(jsonData.(string)).Find(k)
	return
}

// 实现apollo配置中心无缓存HTTP接口GET方法
func (a *appConfig) GET(config_server_url, appId, clusterName, namespaceName, releaseKey, ip string) (j interface{}, err error) {
	// 1、http get apollo配置中心api请求
	url := fmt.Sprintf("http://%s/configs/%s/%s/%s", a.apolloSocket, a.appID, a.cluster, a.namespace)
	httpRet, httpErr := http.Get(url)
	if httpErr != nil {
		err = httpErr
		return
	}
	defer httpRet.Body.Close()
	body, bErr := ioutil.ReadAll(httpRet.Body)
	if bErr != nil {
		err = bErr
		return
	}

	// 2.用jquery解析http返回数据
	dom, gqErr := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	if gqErr != nil {
		err = gqErr
		return
	}
	j = dom.Text()
	return
}
