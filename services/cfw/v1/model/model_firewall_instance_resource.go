package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FirewallInstanceResource struct {

	// 资源id
	ResourceId *string `json:"resource_id,omitempty"`

	// 服务类型，用于CBC使用，特指：hws.service.type.cfw
	CloudServiceType *string `json:"cloud_service_type,omitempty"`

	// 资源类型，包括: 1、云防火墙:hws.resource.type.cfw 2、EIP:hws.resource.type.cfw.exp.eip 3、带宽:hws.resource.type.cfw.exp.bandwidth 4、VPC:hws.resource.type.cfw.exp.vpc 5、日志存储:hws.resource.type.cfw.exp.logaudit
	ResourceType *string `json:"resource_type,omitempty"`

	// 库存单位码
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	// 资源数量
	ResourceSize *int32 `json:"resource_size,omitempty"`

	// 资源单位
	ResourceSizeMeasureId *int32 `json:"resource_size_measure_id,omitempty"`
}

func (o FirewallInstanceResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FirewallInstanceResource struct{}"
	}

	return strings.Join([]string{"FirewallInstanceResource", string(data)}, " ")
}
