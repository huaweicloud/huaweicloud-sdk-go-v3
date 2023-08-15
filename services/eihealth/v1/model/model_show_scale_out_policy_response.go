package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowScaleOutPolicyResponse Response Object
type ShowScaleOutPolicyResponse struct {

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

	// 是否开启cpu执行策略
	CpuRuleEnable *bool `json:"cpu_rule_enable,omitempty"`

	// cpu分配率百分比
	CpuPercent *int32 `json:"cpu_percent,omitempty"`

	// 满足扩容策略中cpu分配率时增加的节点数
	AddNodesForCpuRule *int32 `json:"add_nodes_for_cpu_rule,omitempty"`

	// 是否开启mem执行策略
	MemRuleEnable *bool `json:"mem_rule_enable,omitempty"`

	// mem分配率百分比
	MemPercent *int32 `json:"mem_percent,omitempty"`

	// 满足扩容策略中mem分配率时增加的节点数
	AddNodesForMemRule *int32 `json:"add_nodes_for_mem_rule,omitempty"`
	HttpStatusCode     int    `json:"-"`
}

func (o ShowScaleOutPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScaleOutPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowScaleOutPolicyResponse", string(data)}, " ")
}
