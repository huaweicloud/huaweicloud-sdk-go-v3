package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecAppTaskStatusRequest Request Object
type ShowSecAppTaskStatusRequest struct {

	// 任务ID
	TaskId string `json:"task_id"`
}

func (o ShowSecAppTaskStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecAppTaskStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowSecAppTaskStatusRequest", string(data)}, " ")
}
