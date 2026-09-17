package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueNew **参数解释：** 工作项。 **取值范围：** 不涉及。
type IssueNew struct {

	// **参数解释：** 工作项的更新日期。时间戳格式（示例：1839340800000） 。 **取值范围：** 不涉及。
	UpdatedOn *string `json:"updated_on,omitempty"`

	StoryPoint *StoryPoint `json:"story_point,omitempty"`

	// **参数解释：** 工作项的负责者。 **取值范围：** 不涉及。
	Subject *string `json:"subject,omitempty"`

	Project *Project `json:"project,omitempty"`

	// **参数解释：** 是否有子工作项。 **取值范围：** true（有子工作项） false（没有子工作项）
	IsParent *bool `json:"isParent,omitempty"`

	// **参数解释：** 工作项完成度。 **取值范围：** 不涉及。
	DoneRatio *int32 `json:"done_ratio,omitempty"`

	// **参数解释：** 发布人 。 **取值范围：** 不涉及。
	FindReleaseDev *string `json:"findReleaseDev,omitempty"`

	Tracker *Tracker `json:"tracker,omitempty"`

	// **参数解释：** 工作项列表id。 **取值范围：** 不涉及。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 工作项的开始日期。时间戳格式（示例：1839340800000）。 **取值范围：** 不涉及。
	StartDate *string `json:"start_date,omitempty"`

	AssignedTo *IssueNewAssignedTo `json:"assigned_to,omitempty"`

	StatusAttribute *StatusAttributeVo `json:"status_attribute,omitempty"`

	Severity *Severity `json:"severity,omitempty"`

	// **参数解释：** 工作项发布版本号。 **取值范围：** 不涉及。
	ReleaseDev *string `json:"releaseDev,omitempty"`

	Author *IssueNewAuthor `json:"author,omitempty"`

	// **参数解释：** 工作项的模块。 **取值范围：** 不涉及。
	Module *interface{} `json:"module,omitempty"`

	// **参数解释：** 工作项的截止日期，时间戳格式（示例：1839340800000）。 **取值范围：** 不涉及。
	DueDate *string `json:"due_date,omitempty"`

	// **参数解释：** 工作项的预计工时(单位：人时)。 **取值范围：** 不涉及。
	ExpectedWorkHours *int32 `json:"expected_work_hours,omitempty"`

	Priority *Priority `json:"priority,omitempty"`

	// **参数解释：** 工作项的实际工时（单位：人/时）。 **取值范围：** 不涉及。
	ActualWorkHours *int32 `json:"actual_work_hours,omitempty"`

	// **参数解释：** 是否关注 。 **取值范围：** true（是） false（不是）
	IsWatcher *bool `json:"is_watcher,omitempty"`

	// **参数解释：** 是否删除 。 **取值范围：** true（是） false（不是）
	Deleted *bool `json:"deleted,omitempty"`

	// **参数解释：** 问题解决版本。 **取值范围：** 不涉及。
	FixedVersion *interface{} `json:"fixed_version,omitempty"`

	// **参数解释：** 是否归档。 **取值范围：** true（是） false（不是）
	IsArchived *bool `json:"is_archived,omitempty"`

	// **参数解释：** 工作项的创建时间，时间戳格式（示例：1839340800000）。 **取值范围：** 不涉及。
	CreatedOn *string `json:"created_on,omitempty"`

	// **参数解释：** 工作项的领域 。 **取值范围：** 不涉及。
	Domain *interface{} `json:"domain,omitempty"`

	// **参数解释：** 工作项的开发人员。 **取值范围：** 不涉及。
	Developer *interface{} `json:"developer,omitempty"`

	// **参数解释：** 关闭人员。 **取值范围：** 不涉及。
	Closeder *interface{} `json:"closeder,omitempty"`

	// **参数解释：** 工作项在列表的展示位置 。 **取值范围：** 不涉及。
	Position *string `json:"position,omitempty"`

	// **参数解释：** 关闭标志 。 **取值范围：** 0（打开） 1（关闭）
	ClosedFlag *int32 `json:"closed_flag,omitempty"`

	// **参数解释：** 工作项的抄送人。 **取值范围：** 不涉及。
	AssignedCcUser *string `json:"assigned_cc_user,omitempty"`

	// **参数解释：** 自定义字段。 **取值范围：** 不涉及。
	CustomValueNew *interface{} `json:"custom_value_new,omitempty"`

	Status *Status `json:"status,omitempty"`
}

func (o IssueNew) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueNew struct{}"
	}

	return strings.Join([]string{"IssueNew", string(data)}, " ")
}
