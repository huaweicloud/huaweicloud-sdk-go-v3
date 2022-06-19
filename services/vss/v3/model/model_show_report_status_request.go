package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowReportStatusRequest struct {

	// 任务ID
	TaskId string `json:"task_id"`
}

func (o ShowReportStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReportStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowReportStatusRequest", string(data)}, " ")
}
