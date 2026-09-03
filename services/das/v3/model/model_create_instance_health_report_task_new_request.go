package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceHealthReportTaskNewRequest Request Object
type CreateInstanceHealthReportTaskNewRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *CreateInstanceHealthReportTaskNewRequestBody `json:"body,omitempty"`
}

func (o CreateInstanceHealthReportTaskNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceHealthReportTaskNewRequest struct{}"
	}

	return strings.Join([]string{"CreateInstanceHealthReportTaskNewRequest", string(data)}, " ")
}
