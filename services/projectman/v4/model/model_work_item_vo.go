package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkItemVo **参数解释**： 工作项信息对象，包含工作项ID、标题、编号、分类、状态、责任人等。 **约束限制**： 不涉及。
type WorkItemVo struct {

	// **参数解释**： 工作项唯一ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 工作项标题。 **取值范围**： 不涉及。
	Title *string `json:"title,omitempty"`

	// **参数解释**： 工作项编号。 **取值范围**： 不涉及。
	Number *string `json:"number,omitempty"`

	// **参数解释**： 工作项分类。 **取值范围**： 不涉及。
	Category *string `json:"category,omitempty"`

	Status *StatusVoIpd `json:"status,omitempty"`

	Assignee *UserVo `json:"assignee,omitempty"`

	// **参数解释**： 工作项基线状态。 **取值范围**： - baselined：已基线 - unbaseline：未基线 - \"\"：未基线
	Baseline *string `json:"baseline,omitempty"`

	// **参数解释**： 工作项变更状态。 **取值范围**： - cannot_finish：不可完成
	ChangeStatus *string `json:"change_status,omitempty"`
}

func (o WorkItemVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkItemVo struct{}"
	}

	return strings.Join([]string{"WorkItemVo", string(data)}, " ")
}
