package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHmReportRequest Request Object
type ShowHmReportRequest struct {

	// 任务ID
	TaskId *string `json:"task_id,omitempty"`
}

func (o ShowHmReportRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHmReportRequest struct{}"
	}

	return strings.Join([]string{"ShowHmReportRequest", string(data)}, " ")
}
