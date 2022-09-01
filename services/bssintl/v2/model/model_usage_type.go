package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UsageType struct {

	// 使用量类型编码。如：reqNumber。
	Code *string `json:"code,omitempty" xml:"code"`

	// 使用量类型名称。如：调用次数。
	Name *string `json:"name,omitempty" xml:"name"`

	// 资源类型编码。例如ECS的VM为“hws.resource.type.vm”。
	ResourceTypeCode *string `json:"resource_type_code,omitempty" xml:"resource_type_code"`

	// 云服务类型编码。例如OBS的云服务类型编码为“hws.service.type.obs”。
	ServiceTypeCode *string `json:"service_type_code,omitempty" xml:"service_type_code"`

	// 资源类型名称。例如ECS的资源类型名称为“云主机”。
	ResourceTypeName *string `json:"resource_type_name,omitempty" xml:"resource_type_name"`

	// 云服务类型名称。例如ECS的云服务类型名称为“弹性云服务器”。
	ServiceTypeName *string `json:"service_type_name,omitempty" xml:"service_type_name"`
}

func (o UsageType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UsageType struct{}"
	}

	return strings.Join([]string{"UsageType", string(data)}, " ")
}
