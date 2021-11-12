package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 节点池每次扩容的动作结果记录，用于确定节点池是否还能继续扩容。
type NodePoolCondition struct {
	// Condition类型。

	Type *string `json:"type,omitempty"`
	// Condition当前状态。

	Status *string `json:"status,omitempty"`
	// 上次状态检查时间。

	LastProbeTime *string `json:"lastProbeTime,omitempty"`
	// 上次状态变更时间。

	LastTransitTime *string `json:"lastTransitTime,omitempty"`
	// 上次状态变更原因。

	Reason *string `json:"reason,omitempty"`
	// Condition详细描述。

	Message *string `json:"message,omitempty"`
}

func (o NodePoolCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodePoolCondition struct{}"
	}

	return strings.Join([]string{"NodePoolCondition", string(data)}, " ")
}
