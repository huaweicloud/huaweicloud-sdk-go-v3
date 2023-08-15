package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTasksRulesetsRequest Request Object
type ShowTasksRulesetsRequest struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	// 任务ID
	TaskId string `json:"task_id"`

	// 分页索引，偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的数量,每页最多显示1000条
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowTasksRulesetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTasksRulesetsRequest struct{}"
	}

	return strings.Join([]string{"ShowTasksRulesetsRequest", string(data)}, " ")
}
