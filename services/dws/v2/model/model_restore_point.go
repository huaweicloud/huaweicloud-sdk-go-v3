package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestorePoint **参数解释**： 恢复集群对象信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type RestorePoint struct {

	// **参数解释**： 快照ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	BackRef *string `json:"back_ref,omitempty"`

	// **参数解释**： 恢复时间。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	RestoreTime *int64 `json:"restore_time,omitempty"`

	// **参数解释**： 原始集群ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ClusterId *string `json:"cluster_id,omitempty"`
}

func (o RestorePoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestorePoint struct{}"
	}

	return strings.Join([]string{"RestorePoint", string(data)}, " ")
}
