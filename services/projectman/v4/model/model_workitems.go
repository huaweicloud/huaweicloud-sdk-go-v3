package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Workitems struct {

	// 工作项id
	Id *string `json:"id,omitempty" xml:"id"`

	// 工作项描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 实际工时
	ActualWorkHours *float64 `json:"actual_work_hours,omitempty" xml:"actual_work_hours"`

	AssignedUser *WorkitemUser `json:"assigned_user,omitempty" xml:"assigned_user"`

	Author *WorkitemUser `json:"author,omitempty" xml:"author"`

	// 工作项开始时间
	BeginTime *string `json:"begin_time,omitempty" xml:"begin_time"`

	// 创建时间
	CreatedTime *string `json:"created_time,omitempty" xml:"created_time"`

	// 标签
	Tags *[]WorkitemsTags `json:"tags,omitempty" xml:"tags"`

	Developer *WorkitemUser `json:"developer,omitempty" xml:"developer"`

	// 抄送人
	AssignedCcUser *[]WorkitemUser `json:"assigned_cc_user,omitempty" xml:"assigned_cc_user"`

	// 发现问题的版本
	DiscoverVersion *string `json:"discover_version,omitempty" xml:"discover_version"`

	// 工作项结束时间
	EndTime *string `json:"end_time,omitempty" xml:"end_time"`

	// 工作项进度值
	DoneRatio *string `json:"done_ratio,omitempty" xml:"done_ratio"`

	// 预计工时
	ExpectedWorkHours *float64 `json:"expected_work_hours,omitempty" xml:"expected_work_hours"`

	// 顺序
	Order *string `json:"order,omitempty" xml:"order"`

	// 父工作项的id
	ParentWorkItemId *string `json:"parent_work_item_id,omitempty" xml:"parent_work_item_id"`

	// 发布的版本
	ReleaseVersion *string `json:"release_version,omitempty" xml:"release_version"`

	// 故事点
	StoryPoint *string `json:"story_point,omitempty" xml:"story_point"`

	Domain *WorkitemsDomain `json:"domain,omitempty" xml:"domain"`

	Iteration *WorkitemsIteration `json:"iteration,omitempty" xml:"iteration"`

	Module *WorkitemsModule `json:"module,omitempty" xml:"module"`

	// 工作项优先级
	Priority *string `json:"priority,omitempty" xml:"priority"`

	// 严重的程度 \"提示\", \"一般\", \"严重\", \"致命\"
	Severity *string `json:"severity,omitempty" xml:"severity"`

	Status *WorkitemsStatus `json:"status,omitempty" xml:"status"`

	// 工作项标题
	Subject *string `json:"subject,omitempty" xml:"subject"`

	// 更新时间
	UpdatedTime *string `json:"updated_time,omitempty" xml:"updated_time"`

	// 工作项编号
	Sequence *string `json:"sequence,omitempty" xml:"sequence"`

	// 重要程度 \"关键\", \"重要\", \"一般\", \"提示\"
	Important *string `json:"important,omitempty" xml:"important"`

	// 用户自定义字段
	CustomFields *[]WorkitemCustomField `json:"custom_fields,omitempty" xml:"custom_fields"`
}

func (o Workitems) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Workitems struct{}"
	}

	return strings.Join([]string{"Workitems", string(data)}, " ")
}
