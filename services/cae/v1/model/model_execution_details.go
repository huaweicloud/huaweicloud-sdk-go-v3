package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExecutionDetails struct {

	// 启停策略执行时间。
	LastExecutionTime *string `json:"last_execution_time,omitempty"`

	// 组件启停规则执行结果列表。
	Items *[]ComponentExecutionResult `json:"items,omitempty"`
}

func (o ExecutionDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecutionDetails struct{}"
	}

	return strings.Join([]string{"ExecutionDetails", string(data)}, " ")
}
