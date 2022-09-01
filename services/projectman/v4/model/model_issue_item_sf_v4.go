package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IssueItemSfV4 struct {

	// 实际工时
	ActualWorkHours *float64 `json:"actual_work_hours,omitempty" xml:"actual_work_hours"`

	AssignedUser *IssueUser `json:"assigned_user,omitempty" xml:"assigned_user"`

	Author *IssueUser `json:"author,omitempty" xml:"author"`

	// 工作项开始时间
	BeginTime *int64 `json:"begin_time,omitempty" xml:"begin_time"`

	// 关闭工作项的时间
	ClosedTime *int64 `json:"closed_time,omitempty" xml:"closed_time"`

	// 创建时间
	CreatedTime *int64 `json:"created_time,omitempty" xml:"created_time"`

	// 自定义属性
	CustomFeilds *[]CustomFeildRecord `json:"custom_feilds,omitempty" xml:"custom_feilds"`

	Developer *IssueUser `json:"developer,omitempty" xml:"developer"`

	// 发现问题的版本
	DiscoverVersion *string `json:"discover_version,omitempty" xml:"discover_version"`

	// 工作项结束时间
	EndTime *int64 `json:"end_time,omitempty" xml:"end_time"`

	// 工作项进度值
	DoneRatio *int32 `json:"done_ratio,omitempty" xml:"done_ratio"`

	// 预计工时
	ExpectedWorkHours *float64 `json:"expected_work_hours,omitempty" xml:"expected_work_hours"`

	// 顺序
	Order *int32 `json:"order,omitempty" xml:"order"`

	// 父工作项的id
	ParentIssueId *int32 `json:"parent_issue_id,omitempty" xml:"parent_issue_id"`

	// 发布的版本
	ReleaseVersion *string `json:"release_version,omitempty" xml:"release_version"`

	// 根工作项的id
	RootIssueId *int32 `json:"root_issue_id,omitempty" xml:"root_issue_id"`

	StoryPoint *IssueItemSfV4StoryPoint `json:"story_point,omitempty" xml:"story_point"`

	Domain *IssueItemSfV4Domain `json:"domain,omitempty" xml:"domain"`

	Iteration *IssueItemSfV4Iteration `json:"iteration,omitempty" xml:"iteration"`

	Module *IssueItemSfV4Module `json:"module,omitempty" xml:"module"`

	Priority *IssueItemSfV4Priority `json:"priority,omitempty" xml:"priority"`

	Severity *IssueItemSfV4Severity `json:"severity,omitempty" xml:"severity"`

	Status *IssueItemSfV4Status `json:"status,omitempty" xml:"status"`

	Tracker *IssueItemSfV4Tracker `json:"tracker,omitempty" xml:"tracker"`

	// 工作项标题
	Subject *string `json:"subject,omitempty" xml:"subject"`

	// 工作项更新时间
	UpdatedTime *int64 `json:"updated_time,omitempty" xml:"updated_time"`
}

func (o IssueItemSfV4) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueItemSfV4 struct{}"
	}

	return strings.Join([]string{"IssueItemSfV4", string(data)}, " ")
}
