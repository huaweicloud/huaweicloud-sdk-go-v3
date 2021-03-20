package model

import (
	"encoding/json"

	"strings"
)

// 创建资源空间结构体。
type AddApplication struct {
	// 资源空间名称。

	AppName string `json:"app_name"`
	// 资源空间ID。

	AppId *string `json:"app_id,omitempty"`
	// 迁移前实例ID。

	InstanceId *string `json:"instance_id,omitempty"`
	// 对接的服务名,IoTDA代表华为云设备接入云服务，CTNBGW代表天翼云设备接入服务

	ServiceName *string `json:"service_name,omitempty"`
}

func (o AddApplication) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AddApplication struct{}"
	}

	return strings.Join([]string{"AddApplication", string(data)}, " ")
}
