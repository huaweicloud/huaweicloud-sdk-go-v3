package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeletesResponseResultDeleteIssueDelIssue struct {

	// **参数解释：** 工作项id。 **取值范围：** 不涉及。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 工作项类型。 **取值范围：** 2（任务/Task） 3（缺陷/Bug） 5（Epic） 6（Feature） 7（Story）
	TrackerId *int32 `json:"tracker_id,omitempty"`

	// **参数解释：** 工作项名称 。 **取值范围：** 不涉及。
	Subject *string `json:"subject,omitempty"`

	// **参数解释：** 工作项状态id 。 **取值范围：** 不涉及。
	StatusId *int32 `json:"status_id,omitempty"`

	// **参数解释：** 工作项完成度。 **取值范围：** 不涉及。
	DoneRatio *int32 `json:"done_ratio,omitempty"`

	// **参数解释：** 预计工时(单位：人时)。 **取值范围：** 不涉及。
	ExpectedWorkHours *int32 `json:"expected_work_hours,omitempty"`

	// **参数解释：** 实际工时(单位：人时)。 **取值范围：** 不涉及。
	ActualWorkHours *int32 `json:"actual_work_hours,omitempty"`

	// **参数解释：** 是否完成删除。 **取值范围：** 0（未删除） 1（已删除）
	Deleted *bool `json:"deleted,omitempty"`

	// **参数解释：** 是否归档。 **取值范围：** 0（未归档） 1（已归档）
	IsArchived *bool `json:"is_archived,omitempty"`
}

func (o BatchDeletesResponseResultDeleteIssueDelIssue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeletesResponseResultDeleteIssueDelIssue struct{}"
	}

	return strings.Join([]string{"BatchDeletesResponseResultDeleteIssueDelIssue", string(data)}, " ")
}
