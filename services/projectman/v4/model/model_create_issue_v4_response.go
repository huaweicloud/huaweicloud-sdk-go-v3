package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateIssueV4Response struct {

	// 实际工时
	ActualWorkHours *float64 `json:"actual_work_hours,omitempty"`

	// 抄送人
	AssignedCcUser *[]IssueUser `json:"assigned_cc_user,omitempty"`

	AssignedUser *IssueUser `json:"assigned_user,omitempty"`

	// 开始时间，年-月-日
	BeginTime *string `json:"begin_time,omitempty"`

	Creator *IssueUser `json:"creator,omitempty"`

	// 自定义属性值,不建议使用，建议参考new_custom_fields字段
	CustomFields *[]CustomField `json:"custom_fields,omitempty"`

	// 自定义属性值
	NewCustomFields *[]NewCustomField `json:"new_custom_fields,omitempty"`

	Developer *IssueUser `json:"developer,omitempty"`

	Domain *CreateIssueResponseV4Domain `json:"domain,omitempty"`

	// 工作项进度值
	DoneRatio *int32 `json:"done_ratio,omitempty"`

	// 结束时间，年-月-日
	EndTime *string `json:"end_time,omitempty"`

	// 预计工时
	ExpectedWorkHours *float64 `json:"expected_work_hours,omitempty"`

	// 工作项项id
	Id *int32 `json:"id,omitempty"`

	Project *IssueProjectResponseV4 `json:"project,omitempty"`

	Iteration *IssueItemSfV4Iteration `json:"iteration,omitempty"`

	Module *IssueItemSfV4Module `json:"module,omitempty"`

	ParentIssue *CreateIssueResponseV4ParentIssue `json:"parent_issue,omitempty"`

	Priority *IssueItemSfV4Priority `json:"priority,omitempty"`

	Severity *IssueItemSfV4Severity `json:"severity,omitempty"`

	Status *IssueItemSfV4Status `json:"status,omitempty"`

	Tracker        *CreateIssueResponseV4Tracker `json:"tracker,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o CreateIssueV4Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIssueV4Response struct{}"
	}

	return strings.Join([]string{"CreateIssueV4Response", string(data)}, " ")
}
