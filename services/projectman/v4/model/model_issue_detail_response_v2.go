package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IssueDetailResponseV2 struct {

	// **参数解释：** 工作项的实际工时（单位：人/时）。 **取值范围：** 不涉及。
	ActualWorkHours *float64 `json:"actual_work_hours,omitempty"`

	// **参数解释：** 当前工作项的抄送人。
	AssignedCcUser *[]UserVo `json:"assigned_cc_user,omitempty"`

	AssignedTo *UserVo `json:"assigned_to,omitempty"`

	// **参数解释：** 工作项的预计开始时间，时间戳格式（示例：1754323200000）。 **取值范围：** 不涉及。
	StartDate *string `json:"start_date,omitempty"`

	// **参数解释：** 工作项创建时间，时间戳格式（示例：1754374102000）。 **取值范围：** 不涉及。
	CreatedOn *string `json:"created_on,omitempty"`

	Author *UserVo `json:"author,omitempty"`

	// **参数解释：** 工作项的自定义字段。
	CustomFields *[]CustomFieldV2 `json:"custom_fields,omitempty"`

	CustomValueNew *IssueDetailCustomFieldV2 `json:"custom_value_new,omitempty"`

	Developer *UserVo `json:"developer,omitempty"`

	Domain *IssueDetailResponseV2Domain `json:"domain,omitempty"`

	// **参数解释：** 工作项完成度。 **取值范围：** 不涉及。
	DoneRatio *int32 `json:"done_ratio,omitempty"`

	// **参数解释：** 工作项的预计结束时间，时间戳格式（示例：1754323200000）。 **取值范围：** 不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释：** 工作项的预计完成工时（单位：人/时）。 **取值范围：** 不涉及。
	ExpectedWorkHours *float64 `json:"expected_work_hours,omitempty"`

	// **参数解释：** 工作项id。 **取值范围：** 不涉及。
	Id *int32 `json:"id,omitempty"`

	Project *ProjectVo `json:"project,omitempty"`

	Iteration *IssueDetailResponseV2Iteration `json:"iteration,omitempty"`

	StoryPoint *IssueDetailResponseV2StoryPoint `json:"story_point,omitempty"`

	Module *IssueDetailResponseV2Module `json:"module,omitempty"`

	// **参数解释：** 工作项的标题。 **取值范围：** 不涉及。
	Subject *string `json:"subject,omitempty"`

	ParentIssue *IssueDetailResponseV2ParentIssue `json:"parent_issue,omitempty"`

	Priority *IssueDetailResponseV2Priority `json:"priority,omitempty"`

	Severity *IssueDetailResponseV2Severity `json:"severity,omitempty"`

	Status *IssueDetailResponseV2Status `json:"status,omitempty"`

	// **参数解释：** 工作项发布版本号。 **取值范围：** 不涉及。
	ReleaseDev *string `json:"release_dev,omitempty"`

	// **参数解释：** 缺陷发现版本号（仅Bug类型工作项具备该字段）。 **取值范围：** 不涉及。
	FindReleaseDev *string `json:"find_release_dev,omitempty"`

	Env *IssueDetailResponseV2Env `json:"env,omitempty"`

	Tracker *IssueDetailResponseV2Tracker `json:"tracker,omitempty"`

	// **参数解释：** 工作项的最后更新时间，时间戳格式（示例：1754374102000）。 **取值范围：** 不涉及。
	UpdatedOn *string `json:"updated_on,omitempty"`

	// **参数解释：** 工作项的关闭时间，时间戳格式（示例：1754374102000）。 **取值范围：** 不涉及。
	ClosedTime *string `json:"closed_time,omitempty"`

	// **参数解释：** 工作项描述。 **取值范围：** 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释：** 工作项的附件列表。
	AccessoriesList *[]IssueAccessoryV2 `json:"accessories_list,omitempty"`

	// **参数解释：** 工作项更新的评论内容。 **取值范围：** 不涉及。
	InnerText *string `json:"inner_text,omitempty"`
}

func (o IssueDetailResponseV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueDetailResponseV2 struct{}"
	}

	return strings.Join([]string{"IssueDetailResponseV2", string(data)}, " ")
}
