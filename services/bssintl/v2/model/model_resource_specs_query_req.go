package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceSpecsQueryReq struct {

	// |参数名称：云服务类型编码| |参数的约束及描述：必填，云服务类型编码|
	CloudServiceType string `json:"cloud_service_type"`

	// |参数名称：资源类型编码| |参数的约束及描述：必填，资源类型编码|
	ResourceType string `json:"resource_type"`

	// |参数名称：区域编码| |参数的约束及描述：必填，区域编码|
	RegionCode string `json:"region_code"`

	// |参数名称：计费模式| |参数的约束及描述：必填，1：包年/包月，3：按需|
	ChargeMode string `json:"charge_mode"`

	// |参数名称：过滤条件| |参数的约束及描述：非必填，过滤条件列表，最多1个|
	Filters *[]ResourceSpecsFilter `json:"filters,omitempty"`

	// |参数名称：翻页信息| |参数的约束及描述：非必填，首页查询不携带此参数，非首页查询传入上一页响应返回的next_marker|
	Marker *string `json:"marker,omitempty"`

	// |参数名称：查询条数| |参数的约束及描述：非必填，取值范围1-100，默认值100|
	Limit *int32 `json:"limit,omitempty"`
}

func (o ResourceSpecsQueryReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceSpecsQueryReq struct{}"
	}

	return strings.Join([]string{"ResourceSpecsQueryReq", string(data)}, " ")
}
