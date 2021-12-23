package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceType struct {
	// 资源类型的编码。例如ECS的VM为“hws.resource.type.vm”。

	ResourceTypeCode *string `json:"resource_type_code,omitempty"`
	// 资源类型的名称。

	ResourceTypeName *string `json:"resource_type_name,omitempty"`
	// 资源类型的描述。

	ResourceTypeDesc *string `json:"resource_type_desc,omitempty"`
}

func (o ResourceType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceType struct{}"
	}

	return strings.Join([]string{"ResourceType", string(data)}, " ")
}
