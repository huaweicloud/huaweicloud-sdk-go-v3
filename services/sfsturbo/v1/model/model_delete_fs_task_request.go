package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFsTaskRequest Request Object
type DeleteFsTaskRequest struct {

	// MIME类型, application/json
	ContentType string `json:"Content-Type"`

	// 文件系统id
	ShareId string `json:"share_id"`

	// 任务类型。例，DU任务取值为dir-usage
	Feature string `json:"feature"`

	// 任务id
	TaskId string `json:"task_id"`
}

func (o DeleteFsTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFsTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteFsTaskRequest", string(data)}, " ")
}
