package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecutionActionPolicy 执行动作策略。
type ExecutionActionPolicy struct {

	// 重跑的节点。
	RerunSteps *[]string `json:"rerun_steps,omitempty"`
}

func (o ExecutionActionPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecutionActionPolicy struct{}"
	}

	return strings.Join([]string{"ExecutionActionPolicy", string(data)}, " ")
}
