package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeTaskMetadata **参数解释：** 升级任务元数据 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type UpgradeTaskMetadata struct {

	// **参数解释：** 升级任务ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Uid *string `json:"uid,omitempty"`

	// **参数解释：** 任务创建时间 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// **参数解释：** 任务更新时间 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	UpdateTimestamp *string `json:"updateTimestamp,omitempty"`
}

func (o UpgradeTaskMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeTaskMetadata struct{}"
	}

	return strings.Join([]string{"UpgradeTaskMetadata", string(data)}, " ")
}
