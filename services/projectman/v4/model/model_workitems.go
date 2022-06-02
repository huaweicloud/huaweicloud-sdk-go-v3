package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Workitems struct {

	// 工作项id
	Id *string `json:"id,omitempty"`

	// 工作项描述
	Description *string `json:"description,omitempty"`

	// 实际工时
	ActualWorkHours *float64 `json:"actual_work_hours,omitempty"`

	AssignedUser *WorkitemUser `json:"assigned_user,omitempty"`

	Author *WorkitemUser `json:"author,omitempty"`

	// 工作项开始时间
	BeginTime *string `json:"begin_time,omitempty"`

	// 创建时间
	CreatedTime *string `json:"created_time,omitempty"`

	// 标签
	Tags *[]WorkitemsTags `json:"tags,omitempty"`

	Developer *WorkitemUser `json:"developer,omitempty"`

	// 抄送人
	AssignedCcUser *[]WorkitemUser `json:"assigned_cc_user,omitempty"`

	// 发现问题的版本
	DiscoverVersion *string `json:"discover_version,omitempty"`

	// 工作项结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 工作项进度值
	DoneRatio *string `json:"done_ratio,omitempty"`

	// 预计工时
	ExpectedWorkHours *float64 `json:"expected_work_hours,omitempty"`

	// 顺序
	Order *string `json:"order,omitempty"`

	// 父工作项的id
	ParentWorkItemId *string `json:"parent_work_item_id,omitempty"`

	// 发布的版本
	ReleaseVersion *string `json:"release_version,omitempty"`

	// 故事点
	StoryPoint *string `json:"story_point,omitempty"`

	Domain *WorkitemsDomain `json:"domain,omitempty"`

	Iteration *WorkitemsIteration `json:"iteration,omitempty"`

	Module *WorkitemsModule `json:"module,omitempty"`

	// 工作项优先级
	Priority *string `json:"priority,omitempty"`

	// 严重的程度 \"提示\", \"一般\", \"严重\", \"致命\"
	Severity *string `json:"severity,omitempty"`

	Status *WorkitemsStatus `json:"status,omitempty"`

	// 工作项标题
	Subject *string `json:"subject,omitempty"`

	// 更新时间
	UpdatedTime *string `json:"updated_time,omitempty"`

	// 工作项编号
	Sequence *string `json:"sequence,omitempty"`

	// 重要程度 \"关键\", \"重要\", \"一般\", \"提示\"
	Important *string `json:"important,omitempty"`

	// 用户自定义字段
	CustomFields *[]WorkitemCustomField `json:"custom_fields,omitempty"`
}

func (o Workitems) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Workitems struct{}"
	}

	return strings.Join([]string{"Workitems", string(data)}, " ")
}
