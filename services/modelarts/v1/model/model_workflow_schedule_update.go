package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkflowScheduleUpdate schedule update struct
type WorkflowScheduleUpdate struct {

	// 内容。
	Content map[string]interface{} `json:"content,omitempty"`

	// 使能标志。
	Enable *bool `json:"enable,omitempty"`
}

func (o WorkflowScheduleUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowScheduleUpdate struct{}"
	}

	return strings.Join([]string{"WorkflowScheduleUpdate", string(data)}, " ")
}
