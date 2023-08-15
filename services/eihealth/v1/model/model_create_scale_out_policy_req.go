package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateScaleOutPolicyReq struct {

	// 策略名称
	Name string `json:"name"`

	// 可用区
	AvailabilityZone string `json:"availability_zone"`

	// 规格编码
	SpecCode string `json:"spec_code"`

	// 扩容节点数上限
	MaxNodes int32 `json:"max_nodes"`

	// 扩容节点数下限
	MinNodes int32 `json:"min_nodes"`

	// 额外数据盘规格编码
	DataDiskSpecCode *string `json:"data_disk_spec_code,omitempty"`

	// 额外数据盘大小
	DataDiskSize *int32 `json:"data_disk_size,omitempty"`

	// 是否启用cpu规则
	CpuRuleEnable bool `json:"cpu_rule_enable"`

	// cpu分配率百分比
	CpuPercent int32 `json:"cpu_percent"`

	// 满足扩容策略中cpu分配率时增加的节点数
	AddNodesForCpuRule int32 `json:"add_nodes_for_cpu_rule"`

	// 是否启用mem规则
	MemRuleEnable bool `json:"mem_rule_enable"`

	// mem分配率百分比
	MemPercent int32 `json:"mem_percent"`

	// 满足扩容策略中mem分配率时增加的节点数
	AddNodesForMemRule int32 `json:"add_nodes_for_mem_rule"`
}

func (o CreateScaleOutPolicyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScaleOutPolicyReq struct{}"
	}

	return strings.Join([]string{"CreateScaleOutPolicyReq", string(data)}, " ")
}
