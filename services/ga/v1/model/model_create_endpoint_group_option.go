package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEndpointGroupOption 创建终端节点组的详细信息。
type CreateEndpointGroupOption struct {

	// 终端节点组名称，取值范围：1~64个字符之间，只能由数字、字母、中划线和中文组成。
	Name string `json:"name"`

	// 终端节点组描述信息，取值范围：0~255个字符之间，禁止输入字符：<>。
	Description *string `json:"description,omitempty"`

	// 流量拨分到此组的百分比。
	TrafficDialPercentage *int32 `json:"traffic_dial_percentage,omitempty"`

	// 终端节点组所属区域ID。
	RegionId string `json:"region_id"`

	// 关联监听器列表。
	Listeners []Id `json:"listeners"`
}

func (o CreateEndpointGroupOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEndpointGroupOption struct{}"
	}

	return strings.Join([]string{"CreateEndpointGroupOption", string(data)}, " ")
}
