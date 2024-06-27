package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTaskAssignCasesRequest Request Object
type ListTaskAssignCasesRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 任务uri
	TaskId string `json:"task_id"`

	Body *QueryTaskAssignCasesInfo `json:"body,omitempty"`
}

func (o ListTaskAssignCasesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaskAssignCasesRequest struct{}"
	}

	return strings.Join([]string{"ListTaskAssignCasesRequest", string(data)}, " ")
}
