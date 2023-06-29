package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecAppTaskResultRequest Request Object
type ShowSecAppTaskResultRequest struct {

	// 任务ID
	TaskId string `json:"task_id"`
}

func (o ShowSecAppTaskResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecAppTaskResultRequest struct{}"
	}

	return strings.Join([]string{"ShowSecAppTaskResultRequest", string(data)}, " ")
}
