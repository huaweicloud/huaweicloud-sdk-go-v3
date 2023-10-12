package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowExecutionParamsRequest Request Object
type ShowExecutionParamsRequest struct {

	// 任务id
	TaskId string `json:"task_id"`

	// 执行记录id
	RecordId *string `json:"record_id,omitempty"`
}

func (o ShowExecutionParamsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowExecutionParamsRequest struct{}"
	}

	return strings.Join([]string{"ShowExecutionParamsRequest", string(data)}, " ")
}
