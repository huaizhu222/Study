package main

import (
	"encoding/json"
	"fmt"

	"Rpc.Study.go/nacos_test/config"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func main() {
	serverConfigs := []constant.ServerConfig{
		{
			Port:   8848,
			IpAddr: "127.0.0.1",
		},
	}
	clientConfig := constant.ClientConfig{
		NamespaceId:         "840b2c81-ea09-43a7-b040-55245fc90dbf", // 如果需要支持多namespace，我们可以创建多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
	}
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		fmt.Printf("创建错误")
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "user-web.json",
		Group:  "dev",
	})
	if err != nil {
		panic(err)
	}

	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: "user-web.yaml",
		Group:  "dev",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("配置文件变化")
			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
		},
	})

	if err != nil {
		panic(err)
	}
	fmt.Println(content)

	ServerConfig := config.ServerConfig{}
	json.Unmarshal([]byte(content), &ServerConfig)
	fmt.Println(ServerConfig)
}
