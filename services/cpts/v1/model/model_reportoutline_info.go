package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReportoutlineInfo struct {

	// 平均响应时间
	AvgResponseTime *float64 `json:"avgResponseTime,omitempty" xml:"avgResponseTime"`

	// 分支id
	BranchId *string `json:"branchId,omitempty" xml:"branchId"`

	// 分支名称
	BranchName *string `json:"branchName,omitempty" xml:"branchName"`

	// 用例重试次数
	CaseRetry *float64 `json:"caseRetry,omitempty" xml:"caseRetry"`

	// 已完成的用例数
	CompleteNum *float64 `json:"completeNum,omitempty" xml:"completeNum"`

	// 持续时间
	Duration *float64 `json:"duration,omitempty" xml:"duration"`

	// 结束时间
	EndTime *string `json:"endTime,omitempty" xml:"endTime"`

	// 已执行用例数
	ExecutedNum *float64 `json:"executedNum,omitempty" xml:"executedNum"`

	// 迭代id
	IterationUri *string `json:"iterationUri,omitempty" xml:"iterationUri"`

	// kpi用例数
	KpiCaseCount *float64 `json:"kpiCaseCount,omitempty" xml:"kpiCaseCount"`

	// kpi用例执行次数
	KpiCaseExecuteCount *float64 `json:"kpiCaseExecuteCount,omitempty" xml:"kpiCaseExecuteCount"`

	// kpi用例通过次数
	KpiCasePassCount *float64 `json:"kpiCasePassCount,omitempty" xml:"kpiCasePassCount"`

	// 最大并发数
	MaxUsers *float64 `json:"maxUsers,omitempty" xml:"maxUsers"`

	// 结果为pass的用例数
	PassNum *float64 `json:"passNum,omitempty" xml:"passNum"`

	// 阶段id
	Stage *float64 `json:"stage,omitempty" xml:"stage"`

	// 阶段名称
	StageName *string `json:"stageName,omitempty" xml:"stageName"`

	// 开始时间
	StartTime *string `json:"startTime,omitempty" xml:"startTime"`

	// 成功率
	SuccessRate *float64 `json:"successRate,omitempty" xml:"successRate"`

	// 任务状态
	TaskStatus *float64 `json:"taskStatus,omitempty" xml:"taskStatus"`

	// 总用例数
	TotalNum *float64 `json:"totalNum,omitempty" xml:"totalNum"`

	// 性能tps指标
	Tps *float64 `json:"tps,omitempty" xml:"tps"`

	// 分支uri
	VersionUri *string `json:"versionUri,omitempty" xml:"versionUri"`

	// 工程id
	ProjectId *string `json:"projectId,omitempty" xml:"projectId"`

	// 服务id
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId"`
}

func (o ReportoutlineInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportoutlineInfo struct{}"
	}

	return strings.Join([]string{"ReportoutlineInfo", string(data)}, " ")
}
