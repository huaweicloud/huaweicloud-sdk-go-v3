package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateScaleOutPolicyReq struct {

	// 策略名称
	Name *string `json:"name,omitempty"`

	// 扩容节点数上限
	MaxNodes *int32 `json:"max_nodes,omitempty"`

	// 扩容节点数下限
	MinNodes *int32 `json:"min_nodes,omitempty"`

	// 是否启用cpu规则
	CpuRuleEnable *bool `json:"cpu_rule_enable,omitempty"`

	// cpu分配率百分比
	CpuPercent *int32 `json:"cpu_percent,omitempty"`

	// 满足扩容策略中cpu分配率时增加的节点数
	AddNodesForCpuRule *int32 `json:"add_nodes_for_cpu_rule,omitempty"`

	// 是否启用mem规则
	MemRuleEnable *bool `json:"mem_rule_enable,omitempty"`

	// mem分配率百分比
	MemPercent *int32 `json:"mem_percent,omitempty"`

	// 满足扩容策略中mem分配率时增加的节点数
	AddNodesForMemRule *int32 `json:"add_nodes_for_mem_rule,omitempty"`
}

func (o UpdateScaleOutPolicyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScaleOutPolicyReq struct{}"
	}

	return strings.Join([]string{"UpdateScaleOutPolicyReq", string(data)}, " ")
}
