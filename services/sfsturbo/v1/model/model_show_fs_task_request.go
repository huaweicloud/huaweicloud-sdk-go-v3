package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFsTaskRequest Request Object
type ShowFsTaskRequest struct {

	// MIME类型, application/json
	ContentType string `json:"Content-Type"`

	// 文件系统id
	ShareId string `json:"share_id"`

	// 任务类型。例，DU任务取值为dir-usage
	Feature string `json:"feature"`

	// 任务ID
	TaskId string `json:"task_id"`
}

func (o ShowFsTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFsTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowFsTaskRequest", string(data)}, " ")
}
