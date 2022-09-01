package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 工作项属性
type CreateIssueRequestV4 struct {

	// 实际工时
	ActualWorkHours *float64 `json:"actual_work_hours,omitempty" xml:"actual_work_hours"`

	// 处理人id,对应用户信息的数字id
	AssignedId *int32 `json:"assigned_id,omitempty" xml:"assigned_id"`

	// 开始时间，年-月-日
	BeginTime *string `json:"begin_time,omitempty" xml:"begin_time"`

	// 描述信息
	Description *string `json:"description,omitempty" xml:"description"`

	// 开发者id,对应用户信息的数字id
	DeveloperId *int32 `json:"developer_id,omitempty" xml:"developer_id"`

	// id 领域, 14 '性能', 15 '功能', 16 '可靠性' 17 '网络安全' 18 '可维护性' 19 '其他DFX' 20 '可用性'
	DomainId *int32 `json:"domain_id,omitempty" xml:"domain_id"`

	// 工作项进度值
	DoneRatio *int32 `json:"done_ratio,omitempty" xml:"done_ratio"`

	// 结束时间，年-月-日
	EndTime *string `json:"end_time,omitempty" xml:"end_time"`

	// 预计工时
	ExpectedWorkHours *float64 `json:"expected_work_hours,omitempty" xml:"expected_work_hours"`

	// 迭代id
	IterationId *int32 `json:"iteration_id,omitempty" xml:"iteration_id"`

	// 模块id
	ModuleId *int32 `json:"module_id,omitempty" xml:"module_id"`

	// 标题
	Name string `json:"name" xml:"name"`

	// 父工作项的id,创建子工作项时必填，父工作项的类型tracker_id不能为2,3
	ParentIssueId *int32 `json:"parent_issue_id,omitempty" xml:"parent_issue_id"`

	// 优先级,   1 低,   2 中,   3 高,
	PriorityId int32 `json:"priority_id" xml:"priority_id"`

	// 重要程度,   10 关键,   11 重要,   12 一般,   13 提示,
	SeverityId *int32 `json:"severity_id,omitempty" xml:"severity_id"`

	// 状态   id, 新建   1, 进行中 2, 已解决 3, 测试中 4, 已关闭 5, 已拒绝 6,
	StatusId *int32 `json:"status_id,omitempty" xml:"status_id"`

	// 工作项类型, 2任务/Task,3缺陷/Bug,5Epic,6Feature,7Story;     5 只能为 6 的父工作项类型;     6 只能为 7 的父工作项类型;     7 只能为 2,3的父;
	TrackerId int32 `json:"tracker_id" xml:"tracker_id"`

	// 用户自定义字段
	NewCustomFields *[]NewCustomField `json:"new_custom_fields,omitempty" xml:"new_custom_fields"`
}

func (o CreateIssueRequestV4) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIssueRequestV4 struct{}"
	}

	return strings.Join([]string{"CreateIssueRequestV4", string(data)}, " ")
}
