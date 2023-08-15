package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHistoryRunInfoRequest Request Object
type ShowHistoryRunInfoRequest struct {

	// 任务id
	TaskId int32 `json:"task_id"`
}

func (o ShowHistoryRunInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHistoryRunInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowHistoryRunInfoRequest", string(data)}, " ")
}
