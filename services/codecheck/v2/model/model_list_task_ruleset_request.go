package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListTaskRulesetRequest struct {
	// 设置媒体类型和编码格式

	ContentType string `json:"Content-Type"`
	// 任务ID

	TaskId string `json:"task_id"`
}

func (o ListTaskRulesetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaskRulesetRequest struct{}"
	}

	return strings.Join([]string{"ListTaskRulesetRequest", string(data)}, " ")
}
