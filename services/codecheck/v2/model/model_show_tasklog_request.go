package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowTasklogRequest struct {

	// 项目ID
	ProjectId string `json:"project_id" xml:"project_id"`

	// 任务ID
	TaskId string `json:"task_id" xml:"task_id"`

	// 任务单次的执行ID
	ExecuteId *string `json:"execute_id,omitempty" xml:"execute_id"`
}

func (o ShowTasklogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTasklogRequest struct{}"
	}

	return strings.Join([]string{"ShowTasklogRequest", string(data)}, " ")
}
