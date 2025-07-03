package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetTaskResultRequest Request Object
type SetTaskResultRequest struct {

	// 项目id
	ProjectUuid string `json:"project_uuid"`

	// 任务uri
	TaskUri string `json:"task_uri"`

	Body *ExecuteTaskInfo `json:"body,omitempty"`
}

func (o SetTaskResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetTaskResultRequest struct{}"
	}

	return strings.Join([]string{"SetTaskResultRequest", string(data)}, " ")
}
