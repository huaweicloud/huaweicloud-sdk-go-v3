package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryDiagnosisResult 一键诊断响应体。
type QueryDiagnosisResult struct {

	// 得分。
	Score *int32 `json:"score,omitempty"`

	// 状态。
	Status *string `json:"status,omitempty"`

	// 进度。
	Progress *int32 `json:"progress,omitempty"`

	// 失败原因。
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 开始时间。
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间。
	EndTime *string `json:"end_time,omitempty"`

	// 诊断项总数。
	Total *int64 `json:"total,omitempty"`

	// 正常数量。
	NormalCount *int64 `json:"normal_count,omitempty"`

	// 异常数量。
	AbnormalCount *int64 `json:"abnormal_count,omitempty"`

	// 告警数量。
	AlarmCount *int64 `json:"alarm_count,omitempty"`

	// 失败数量。
	FailureCount *int64 `json:"failure_count,omitempty"`

	// 超时数量。
	TimeoutCount *int64 `json:"timeout_count,omitempty"`

	// 诊断结果。
	DiagnosisResults *[]QueryDiagnosisResultDiagnosisResults `json:"diagnosis_results,omitempty"`
}

func (o QueryDiagnosisResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryDiagnosisResult struct{}"
	}

	return strings.Join([]string{"QueryDiagnosisResult", string(data)}, " ")
}
