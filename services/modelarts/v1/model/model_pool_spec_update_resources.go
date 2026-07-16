package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PoolSpecUpdateResources struct {

	// **参数解释**：资源规格。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Flavor string `json:"flavor"`

	// **参数解释**：相应规格的资源数量。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Count int32 `json:"count"`

	// **参数解释**：更新的AZ列表。
	Azs *[]PoolNodeAz `json:"azs,omitempty"`
}

func (o PoolSpecUpdateResources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolSpecUpdateResources struct{}"
	}

	return strings.Join([]string{"PoolSpecUpdateResources", string(data)}, " ")
}
