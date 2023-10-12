package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Listener 监听器
type Listener struct {

	// 监听器名称
	Name *string `json:"name,omitempty"`

	// 监听器对外提供服务端口
	Port *int32 `json:"port,omitempty"`

	Backend *Backend `json:"backend,omitempty"`

	// 创建负载均衡器的IP协议类型
	IpVersion *int32 `json:"ip_version,omitempty"`
}

func (o Listener) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Listener struct{}"
	}

	return strings.Join([]string{"Listener", string(data)}, " ")
}
