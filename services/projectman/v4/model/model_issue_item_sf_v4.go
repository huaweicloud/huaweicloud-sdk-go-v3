package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IssueItemSfV4 struct {

	// 实际工时
	ActualWorkHours *float64 `json:"actual_work_hours,omitempty"`

	AssignedUser *IssueUser `json:"assigned_user,omitempty"`

	Author *IssueUser `json:"author,omitempty"`

	// 工作项开始时间
	BeginTime *int64 `json:"begin_time,omitempty"`

	// 关闭工作项的时间
	ClosedTime *int64 `json:"closed_time,omitempty"`

	// 创建时间
	CreatedTime *int64 `json:"created_time,omitempty"`

	// 自定义属性
	CustomFeilds *[]CustomFeildRecord `json:"custom_feilds,omitempty"`

	Developer *IssueUser `json:"developer,omitempty"`

	// 发现问题的版本
	DiscoverVersion *string `json:"discover_version,omitempty"`

	// 工作项结束时间
	EndTime *int64 `json:"end_time,omitempty"`

	// 工作项进度值
	DoneRatio *int32 `json:"done_ratio,omitempty"`

	// 预计工时
	ExpectedWorkHours *float64 `json:"expected_work_hours,omitempty"`

	// 顺序
	Order *int32 `json:"order,omitempty"`

	// 父工作项的id
	ParentIssueId *int32 `json:"parent_issue_id,omitempty"`

	// 发布的版本
	ReleaseVersion *string `json:"release_version,omitempty"`

	// 根工作项的id
	RootIssueId *int32 `json:"root_issue_id,omitempty"`

	StoryPoint *IssueItemSfV4StoryPoint `json:"story_point,omitempty"`

	Domain *IssueItemSfV4Domain `json:"domain,omitempty"`

	Iteration *IssueItemSfV4Iteration `json:"iteration,omitempty"`

	Module *IssueItemSfV4Module `json:"module,omitempty"`

	Priority *IssueItemSfV4Priority `json:"priority,omitempty"`

	Severity *IssueItemSfV4Severity `json:"severity,omitempty"`

	Status *IssueItemSfV4Status `json:"status,omitempty"`

	Tracker *IssueItemSfV4Tracker `json:"tracker,omitempty"`

	// 工作项标题
	Subject *string `json:"subject,omitempty"`

	// 工作项更新时间
	UpdatedTime *int64 `json:"updated_time,omitempty"`
}

func (o IssueItemSfV4) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueItemSfV4 struct{}"
	}

	return strings.Join([]string{"IssueItemSfV4", string(data)}, " ")
}
