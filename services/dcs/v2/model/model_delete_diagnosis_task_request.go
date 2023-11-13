package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDiagnosisTaskRequest Request Object
type DeleteDiagnosisTaskRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *DeleteDiagnosisReportRequest `json:"body,omitempty"`
}

func (o DeleteDiagnosisTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDiagnosisTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteDiagnosisTaskRequest", string(data)}, " ")
}
