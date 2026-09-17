package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SnapshotCluserResponseMetadata **参数解释：** 备份任务数据 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type SnapshotCluserResponseMetadata struct {

	// **参数解释：** API版本，默认为v3.1 **约束限制：** 不涉及 **取值范围：** - v3.1  **默认取值：** v3.1
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数解释：** 任务类型 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Kind *string `json:"kind,omitempty"`
}

func (o SnapshotCluserResponseMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SnapshotCluserResponseMetadata struct{}"
	}

	return strings.Join([]string{"SnapshotCluserResponseMetadata", string(data)}, " ")
}
