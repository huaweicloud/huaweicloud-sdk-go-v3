package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateTaskRulesetRequest struct {
	// 设置媒体类型和编码格式

	ContentType string `json:"Content-Type"`
	// 任务ID

	TaskId string `json:"task_id"`
	// 修改任务规则集

	Body *[]UpdateTaskRulesetItem `json:"body,omitempty"`
}

func (o UpdateTaskRulesetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskRulesetRequest struct{}"
	}

	return strings.Join([]string{"UpdateTaskRulesetRequest", string(data)}, " ")
}
