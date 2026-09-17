package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Volume **参数解释**： 磁盘信息对象。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type Volume struct {

	// **参数解释**： 磁盘名称。 **约束限制**： 不涉及。 **取值范围**： - SSD：超高IO - SAS：高IO - SATA：普通IO  **默认取值**： 不涉及。
	Volume string `json:"volume"`

	// **参数解释**： 磁盘容量，单位：GB。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Capacity *int32 `json:"capacity,omitempty"`

	// **参数解释**： ESSD2磁盘IOPS值。  **约束限制**： 仅磁盘类型为ESSD2时必填。 **取值范围**： 100~256000 **默认取值**： 不涉及。
	Iops *int32 `json:"iops,omitempty"`
}

func (o Volume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Volume struct{}"
	}

	return strings.Join([]string{"Volume", string(data)}, " ")
}
