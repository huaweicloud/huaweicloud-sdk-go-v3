package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PrecheckTaskMetadata **参数解释：** 升级前检查任务元数据 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type PrecheckTaskMetadata struct {

	// **参数解释：** 任务ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Uid *string `json:"uid,omitempty"`

	// **参数解释：** 任务创建时间 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// **参数解释：** 任务更新时间 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	UpdateTimestamp *string `json:"updateTimestamp,omitempty"`
}

func (o PrecheckTaskMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrecheckTaskMetadata struct{}"
	}

	return strings.Join([]string{"PrecheckTaskMetadata", string(data)}, " ")
}
