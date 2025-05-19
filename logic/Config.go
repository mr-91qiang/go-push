package logic

import (
	"encoding/json"
	"io/ioutil"
)

type GatewayConfig struct {
	Hostname string `json:"hostname"`
	Port     int    `json:"port"`
}

// 程序配置
type Config struct {
	ServicePort                int             `json:"servicePort"`                // 服务端口
	ServiceReadTimeout         int             `json:"serviceReadTimeout"`         // 接口读超时
	ServiceWriteTimeout        int             `json:"serviceWriteTimeout"`        // 接口写超时
	GatewayList                []GatewayConfig `json:"gatewayList"`                // 网关列表
	GatewayMaxConnection       int             `json:"gatewayMaxConnection"`       // 每个网关的最多并发连接数
	GatewayTimeout             int             `json:"gatewayTimeout"`             // 网关单个请求的超时时间
	GatewayIdleTimeout         int             `json:"gatewayIdleTimeout"`         // 网关连接的空闲关闭时间
	GatewayDispatchWorkerCount int             `json:"gatewayDispatchWorkerCount"` // 向各个网关分发消息的协程数量
	GatewayDispatchChannelSize int             `json:"gatewayDispatchChannelSize"` // 待分发消息队列长度
	GatewayMaxPendingCount     int             `json:"gatewayMaxPendingCount"`     // 网关单个请求的超时时间
	GatewayPushRetry           int             `json:"gatewayPushRetry"`           // 网关推送重试次数
}

var (
	G_config *Config
)

// 初始化配置文件
func InitConfig(filename string) (err error) {
	var (
		content []byte
		conf    Config
	)

	if content, err = ioutil.ReadFile(filename); err != nil {
		return
	}

	if err = json.Unmarshal(content, &conf); err != nil {
		return
	}

	G_config = &conf
	return
}
