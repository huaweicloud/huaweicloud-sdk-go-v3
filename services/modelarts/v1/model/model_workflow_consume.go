package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkflowConsume workflow consume struct
type WorkflowConsume struct {

	// 付费工作流计费周期。
	Period *string `json:"period,omitempty"`

	// 付费工作流可使用的时间值。
	Value *int64 `json:"value,omitempty"`
}

func (o WorkflowConsume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowConsume struct{}"
	}

	return strings.Join([]string{"WorkflowConsume", string(data)}, " ")
}
