package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowTaskDefectsStatisticRequest struct {

	// 任务ID
	TaskId string `json:"task_id"`
}

func (o ShowTaskDefectsStatisticRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskDefectsStatisticRequest struct{}"
	}

	return strings.Join([]string{"ShowTaskDefectsStatisticRequest", string(data)}, " ")
}
