package main

import (
	apolloConf "ClessLi/apollo/config"
	"flag"
	"fmt"
	"github.com/robfig/config"
	"os"
)

// 定义程序入参
var iniConfPath *string = flag.String("f", "/home/shell/docker/docker.ini", "`Configfile` storage path.")
var jobName *string = flag.String("j", "", "`JobName` for the project of jenkins.")
var env *string = flag.String("e", "", "Name of task execution `env`ironment.")

func main() {
	flag.Parse()
	// 判断docker.ini配置文件是否存在
	isExist, pathErr := PathExists(*iniConfPath)
	if !isExist {
		if pathErr != nil {
			fmt.Println("The config", *iniConfPath, "is not found.")
		} else {
			fmt.Println("Unkown error of the configfile.")
		}
		flag.Usage()
		os.Exit(1)
	}

	// 任务名和环境为必选
	if *jobName == "" || *env == "" {
		//Usage()
		flag.Usage()
		os.Exit(1)
	}

	// 读取docker.ini配置
	conf, confErr := config.ReadDefault(*iniConfPath)
	if confErr != nil {
		fmt.Println(confErr)
		flag.Usage()
		os.Exit(1)
	}

	// 比对目标任务名
	appId, _ := conf.String(*jobName, "APOLLO_APP_ID")
	if appId == "" {
		fmt.Println("The project", *jobName, "is not found at", *env, ".")
		flag.Usage()
		os.Exit(1)
	}
	// 读取是否有指定的apollo集群名、命名空间、dubbo服务指定注册端口配置的键名
	cluster, _ := conf.String(*jobName, "APOLLO_CLUSTER")
	ns, _ := conf.String(*jobName, "APOLLO_NAMESPACE")
	dubboPortKey, _ := conf.String(*jobName, "APOLLO_DUBBO_PORT_KEY")

	// 获取dubbo服务指定注册端口
	port, portErr := apolloConf.GetDubboPort(*env, appId, cluster, ns, dubboPortKey)
	if portErr != nil {
		fmt.Println(pathErr)
		flag.Usage()
		os.Exit(1)
	}
	fmt.Println(port)
}

// 文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, err
	} else {
		return false, nil
	}
}
