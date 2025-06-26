package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DiagnosisTaskNodeDetail 诊断步骤结果对象
type DiagnosisTaskNodeDetail struct {

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

	// 诊断步骤主键ID
	DiagnosisRecordId *string `json:"diagnosis_record_id,omitempty"`

	// 诊断步骤开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 诊断步骤结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 诊断步骤执行结果
	Message *string `json:"message,omitempty"`
}

func (o DiagnosisTaskNodeDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiagnosisTaskNodeDetail struct{}"
	}

	return strings.Join([]string{"DiagnosisTaskNodeDetail", string(data)}, " ")
}
