package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DiagnosisTaskNode 单诊断步骤结果对象
type DiagnosisTaskNode struct {

	// id
	Id *string `json:"id,omitempty"`

	// code
	Code *string `json:"code,omitempty"`

	// 诊断步骤名称
	Name *string `json:"name,omitempty"`

	// 诊断步骤名称(中文)
	NameZh *string `json:"name_zh,omitempty"`

	// 诊断任务ID
	DiagnosisTaskId *string `json:"diagnosis_task_id,omitempty"`

	// 任务状态，waiting,executing,failed,finish,cancel
	Status *string `json:"status,omitempty"`
}

func (o DiagnosisTaskNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiagnosisTaskNode struct{}"
	}

	return strings.Join([]string{"DiagnosisTaskNode", string(data)}, " ")
}
