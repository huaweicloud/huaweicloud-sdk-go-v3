package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCpiTaskResultRequest Request Object
type ShowCpiTaskResultRequest struct {

	// CPI任务ID
	TaskId string `json:"task_id"`
}

func (o ShowCpiTaskResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCpiTaskResultRequest struct{}"
	}

	return strings.Join([]string{"ShowCpiTaskResultRequest", string(data)}, " ")
}
