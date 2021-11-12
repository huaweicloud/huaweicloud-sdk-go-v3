package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UsageType struct {
	// 使用量类型编码。如：reqNumber。

	Code *string `json:"code,omitempty"`
	// 使用量类型名称。如：调用次数。

	Name *string `json:"name,omitempty"`
	// 资源类型编码。例如ECS的VM为“hws.resource.type.vm”。

	ResourceTypeCode *string `json:"resource_type_code,omitempty"`
	// 云服务类型编码。例如ECS的云服务类型编码为“hws.service.type.ec2”。

	ServiceTypeCode *string `json:"service_type_code,omitempty"`
}

func (o UsageType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UsageType struct{}"
	}

	return strings.Join([]string{"UsageType", string(data)}, " ")
}
