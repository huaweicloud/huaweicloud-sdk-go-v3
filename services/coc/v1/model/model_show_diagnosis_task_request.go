package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDiagnosisTaskRequest Request Object
type ShowDiagnosisTaskRequest struct {

	// 诊断工单ID
	TaskId string `json:"task_id"`

	// 非必填，被诊断实例的ID，传入后查询该实例诊断报告
	InstanceId string `json:"instance_id"`
}

func (o ShowDiagnosisTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDiagnosisTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowDiagnosisTaskRequest", string(data)}, " ")
}
