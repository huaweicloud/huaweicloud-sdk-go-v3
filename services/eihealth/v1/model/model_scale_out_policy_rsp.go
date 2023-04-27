package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ScaleOutPolicyRsp struct {

	// 策略ID
	Id *string `json:"id,omitempty"`

	// 策略名称
	Name *string `json:"name,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	NodeSpec *NodeSpecDto `json:"node_spec,omitempty"`

	// 可用区
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 节点数量
	Nodes *int32 `json:"nodes,omitempty"`

	// 扩容节点数上限
	MaxNodes *int32 `json:"max_nodes,omitempty"`

	// 扩容节点数下限
	MinNodes *int32 `json:"min_nodes,omitempty"`

	// 伸缩次数
	ScalingTimes *int32 `json:"scaling_times,omitempty"`

	// 是否开启自动扩容
	ScalingEnable *bool `json:"scaling_enable,omitempty"`
}

func (o ScaleOutPolicyRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScaleOutPolicyRsp struct{}"
	}

	return strings.Join([]string{"ScaleOutPolicyRsp", string(data)}, " ")
}
