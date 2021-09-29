package model

import (
	"encoding/json"

	"strings"
)

type ReportoutlineInfo struct {
	// avgResponseTime

	AvgResponseTime *float32 `json:"avgResponseTime,omitempty"`
	// caseRetry

	CaseRetry *int32 `json:"caseRetry,omitempty"`
	// completeNum

	CompleteNum *int32 `json:"completeNum,omitempty"`
	// duration

	Duration *int32 `json:"duration,omitempty"`
	// endTime

	EndTime *string `json:"endTime,omitempty"`
	// executedNum

	ExecutedNum *int32 `json:"executedNum,omitempty"`
	// iterationUri

	IterationUri *string `json:"iterationUri,omitempty"`
	// kpiCaseCount

	KpiCaseCount *int32 `json:"kpiCaseCount,omitempty"`
	// kpiCaseExecuteCount

	KpiCaseExecuteCount *int32 `json:"kpiCaseExecuteCount,omitempty"`
	// kpiCasePassCount

	KpiCasePassCount *int32 `json:"kpiCasePassCount,omitempty"`
	// maxUsers

	MaxUsers *int32 `json:"maxUsers,omitempty"`
	// passNum

	PassNum *int32 `json:"passNum,omitempty"`
	// progressState

	ProgressState *string `json:"progressState,omitempty"`
	// stage

	Stage *int32 `json:"stage,omitempty"`
	// stageName

	StageName *string `json:"stageName,omitempty"`
	// startTime

	StartTime *string `json:"startTime,omitempty"`
	// statusValue

	StatusValue *string `json:"statusValue,omitempty"`
	// successRate

	SuccessRate *int32 `json:"successRate,omitempty"`
	// taskStatus

	TaskStatus *int32 `json:"taskStatus,omitempty"`
	// totalNum

	TotalNum *int32 `json:"totalNum,omitempty"`
	// tps

	Tps *float32 `json:"tps,omitempty"`
	// versionUri

	VersionUri *string `json:"versionUri,omitempty"`
}

func (o ReportoutlineInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ReportoutlineInfo struct{}"
	}

	return strings.Join([]string{"ReportoutlineInfo", string(data)}, " ")
}
