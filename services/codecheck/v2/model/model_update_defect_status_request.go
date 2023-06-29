package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDefectStatusRequest Request Object
type UpdateDefectStatusRequest struct {

	// 任务id
	TaskId string `json:"task_id"`

	Body *UpdateDefectRequestBody `json:"body,omitempty"`
}

func (o UpdateDefectStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDefectStatusRequest struct{}"
	}

	return strings.Join([]string{"UpdateDefectStatusRequest", string(data)}, " ")
}
