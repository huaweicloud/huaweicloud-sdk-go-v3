package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceTypes struct {

	// 资源类型的编码。例如ECS的VM为“hws.resource.type.vm”。
	ResourceTypeCode *string `json:"resource_type_code,omitempty"`

	// 资源类型的名称。
	ResourceTypeName *string `json:"resource_type_name,omitempty"`

	// 资源类型的描述。
	ResourceTypeDesc *string `json:"resource_type_desc,omitempty"`

	// 资源类型归属的服务类型编码。例如：hws.service.type.offline。
	ServiceTypeCode *string `json:"service_type_code,omitempty"`
}

func (o ResourceTypes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTypes struct{}"
	}

	return strings.Join([]string{"ResourceTypes", string(data)}, " ")
}
