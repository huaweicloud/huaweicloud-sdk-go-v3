package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新终端节点组的详细信息。
type UpdateEndpointGroupOption struct {

	// 终端节点组名称，取值范围：1~64个字符之间，只能由数字、字母、中划线和中文组成。
	Name *string `json:"name,omitempty"`

	// 终端节点组描述信息，取值范围：0~255个字符之间，禁止输入字符：<>。
	Description *string `json:"description,omitempty"`

	// 流量拨分到此组的百分比。
	TrafficDialPercentage *int32 `json:"traffic_dial_percentage,omitempty"`
}

func (o UpdateEndpointGroupOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointGroupOption struct{}"
	}

	return strings.Join([]string{"UpdateEndpointGroupOption", string(data)}, " ")
}
