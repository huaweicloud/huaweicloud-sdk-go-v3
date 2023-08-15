package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueRequestV4 工作项属性
type IssueRequestV4 struct {

	// 实际工时
	ActualWorkHours *float64 `json:"actual_work_hours,omitempty"`

	// 处理人id,对应用户信息的数字id
	AssignedId *int32 `json:"assigned_id,omitempty"`

	// 开始时间，年-月-日
	BeginTime *string `json:"begin_time,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 开发者id,对应用户信息的数字id
	DeveloperId *int32 `json:"developer_id,omitempty"`

	// 领域id
	DomainId *int32 `json:"domain_id,omitempty"`

	// 工作项进度值
	DoneRatio *int32 `json:"done_ratio,omitempty"`

	// 结束时间，年-月-日
	EndTime *string `json:"end_time,omitempty"`

	// 预计工时
	ExpectedWorkHours *float64 `json:"expected_work_hours,omitempty"`

	// 迭代id
	IterationId *int32 `json:"iteration_id,omitempty"`

	// 模块id
	ModuleId *int32 `json:"module_id,omitempty"`

	// 标题
	Name *string `json:"name,omitempty"`

	// 父工作项的id
	ParentIssueId *int32 `json:"parent_issue_id,omitempty"`

	// 优先级,   1 低,   2 中,   3 高,
	PriorityId *int32 `json:"priority_id,omitempty"`

	// 重要程度,   10 关键,   11 重要,   12 一般,   13 提示,
	SeverityId *int32 `json:"severity_id,omitempty"`

	// 状态   id, 新建   1, 进行中 2, 已解决 3, 测试中 4, 已关闭 5, 已解决 6,
	StatusId *int32 `json:"status_id,omitempty"`

	// 工作项类型,2任务/Task,3缺陷/Bug,5Epic,6Feature,7Story
	TrackerId *int32 `json:"tracker_id,omitempty"`

	// 用户自定义字段
	NewCustomFields *[]NewCustomField `json:"new_custom_fields,omitempty"`
}

func (o IssueRequestV4) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueRequestV4 struct{}"
	}

	return strings.Join([]string{"IssueRequestV4", string(data)}, " ")
}
