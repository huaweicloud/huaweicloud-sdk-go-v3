package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelInstanceScheduleWindowRequest Request Object
type CancelInstanceScheduleWindowRequest struct {

	// 任务ID，取值为定时任务列表返回job_id字段。
	JobId string `json:"job_id"`
}

func (o CancelInstanceScheduleWindowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelInstanceScheduleWindowRequest struct{}"
	}

	return strings.Join([]string{"CancelInstanceScheduleWindowRequest", string(data)}, " ")
}
