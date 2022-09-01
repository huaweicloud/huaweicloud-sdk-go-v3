package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateTaskRulesetRequest struct {

	// 任务ID
	TaskId string `json:"task_id" xml:"task_id"`

	// 修改任务规则集
	Body *[]UpdateTaskRulesetItem `json:"body,omitempty" xml:"body"`
}

func (o UpdateTaskRulesetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskRulesetRequest struct{}"
	}

	return strings.Join([]string{"UpdateTaskRulesetRequest", string(data)}, " ")
}
