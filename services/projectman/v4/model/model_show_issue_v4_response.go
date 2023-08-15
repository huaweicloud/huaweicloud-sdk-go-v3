package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIssueV4Response Response Object
type ShowIssueV4Response struct {

	// 实际工时
	ActualWorkHours *float64 `json:"actual_work_hours,omitempty"`

	// 抄送人
	AssignedCcUser *[]IssueUser `json:"assigned_cc_user,omitempty"`

	AssignedUser *IssueUser `json:"assigned_user,omitempty"`

	// 预计开始时间，年-月-日
	BeginTime *string `json:"begin_time,omitempty"`

	// 创建时间 年-月-日 时:分:秒
	CreatedTime *string `json:"created_time,omitempty"`

	Creator *IssueUser `json:"creator,omitempty"`

	// 自定义属性值,不建议使用，建议参考new_custom_fields字段
	CustomFields *[]CustomField `json:"custom_fields,omitempty"`

	// 自定义属性值
	NewCustomFields *[]IssueDetailCustomField `json:"new_custom_fields,omitempty"`

	Developer *IssueUser `json:"developer,omitempty"`

	Domain *IssueItemSfV4Domain `json:"domain,omitempty"`

	// 工作项进度值
	DoneRatio *int32 `json:"done_ratio,omitempty"`

	// 预计结束时间，年-月-日
	EndTime *string `json:"end_time,omitempty"`

	// 预计工时
	ExpectedWorkHours *float64 `json:"expected_work_hours,omitempty"`

	// 工作项项id
	Id *int32 `json:"id,omitempty"`

	Project *IssueProjectResponseV4 `json:"project,omitempty"`

	Iteration *IssueItemSfV4Iteration `json:"iteration,omitempty"`

	StoryPoint *IssueDetailResponseV4StoryPoint `json:"story_point,omitempty"`

	Module *IssueItemSfV4Module `json:"module,omitempty"`

	// 标题
	Name *string `json:"name,omitempty"`

	ParentIssue *CreateIssueResponseV4ParentIssue `json:"parent_issue,omitempty"`

	Priority *IssueItemSfV4Priority `json:"priority,omitempty"`

	Severity *IssueItemSfV4Severity `json:"severity,omitempty"`

	Status *IssueItemSfV4Status `json:"status,omitempty"`

	// 工作项发布版本号
	ReleaseDev *string `json:"release_dev,omitempty"`

	// 缺陷发现版本号（仅Bug类型工作项具备该字段）
	FindReleaseDev *string `json:"find_release_dev,omitempty"`

	Env *IssueDetailResponseV4Env `json:"env,omitempty"`

	Tracker *CreateIssueResponseV4Tracker `json:"tracker,omitempty"`

	// 更新时间 年-月-日 时:分:秒
	UpdatedTime *string `json:"updated_time,omitempty"`

	// 关闭时间 年-月-日 时:分:秒
	ClosedTime *string `json:"closed_time,omitempty"`

	// 工作项描述
	Description *string `json:"description,omitempty"`

	Order *IssueOrder `json:"order,omitempty"`

	// 附近列表
	Accessories    *[]IssueAccessory `json:"accessories,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowIssueV4Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIssueV4Response struct{}"
	}

	return strings.Join([]string{"ShowIssueV4Response", string(data)}, " ")
}
