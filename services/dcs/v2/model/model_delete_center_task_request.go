package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCenterTaskRequest Request Object
type DeleteCenterTaskRequest struct {

	// 后台任务ID
	TaskId string `json:"task_id"`

	Body *DeleteCenterTaskRequestBody `json:"body,omitempty"`
}

func (o DeleteCenterTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCenterTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteCenterTaskRequest", string(data)}, " ")
}
