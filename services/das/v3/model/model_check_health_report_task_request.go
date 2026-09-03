package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckHealthReportTaskRequest Request Object
type CheckHealthReportTaskRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o CheckHealthReportTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckHealthReportTaskRequest struct{}"
	}

	return strings.Join([]string{"CheckHealthReportTaskRequest", string(data)}, " ")
}
