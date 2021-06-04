package model

import (
	"encoding/json"

	"strings"
)

// 创建资源空间结构体。
type AddApplication struct {
	// **参数说明**：资源空间名称。 **取值范围**：长度不超过64，只允许字母、数字、下划线（_）、连接符（-）的组合。

	AppName string `json:"app_name"`
	// **参数说明**：资源空间ID。 **取值范围**：长度不超过64，只允许字母、数字、下划线（_）、连接符（-）的组合。

	AppId *string `json:"app_id,omitempty"`
	// **参数说明**：迁移前实例ID。 **取值范围**：长度不超过64，只允许字母、数字、下划线（_）、连接符（-）的组合。

	InstanceId *string `json:"instance_id,omitempty"`
	// **参数说明**：对接的服务名。 **取值范围**： - IoTDA：代表华为云设备接入云服务。 - CTNBGW：代表天翼云设备接入服务。

	ServiceName *string `json:"service_name,omitempty"`
}

func (o AddApplication) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AddApplication struct{}"
	}

	return strings.Join([]string{"AddApplication", string(data)}, " ")
}
