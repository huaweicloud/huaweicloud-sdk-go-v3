package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CloudServiceBasic struct {

	// |参数名称：云服务类型编码| |参数的约束及描述：云服务类型编码|
	CloudServiceType *string `json:"cloud_service_type,omitempty"`

	// |参数名称：云服务类型名称| |参数的约束及描述：云服务类型名称|
	CloudServiceTypeName *string `json:"cloud_service_type_name,omitempty"`

	// |参数名称：资源类型编码| |参数的约束及描述：资源类型编码|
	ResourceType *string `json:"resource_type,omitempty"`

	// |参数名称：资源类型名称| |参数的约束及描述：资源类型名称|
	ResourceTypeName *string `json:"resource_type_name,omitempty"`

	// |参数名称：资源规格编码| |参数的约束及描述：资源规格编码|
	ResourceSpec *string `json:"resource_spec,omitempty"`

	// |参数名称：资源规格名称| |参数的约束及描述：资源规格名称|
	ResourceSpecName *string `json:"resource_spec_name,omitempty"`

	// |参数名称：计费模式| |参数的约束及描述：1：包年/包月，3：按需|
	ChargeMode *string `json:"charge_mode,omitempty"`

	// |参数名称：区域编码| |参数的约束及描述：区域编码|
	RegionCode *string `json:"region_code,omitempty"`
}

func (o CloudServiceBasic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloudServiceBasic struct{}"
	}

	return strings.Join([]string{"CloudServiceBasic", string(data)}, " ")
}
