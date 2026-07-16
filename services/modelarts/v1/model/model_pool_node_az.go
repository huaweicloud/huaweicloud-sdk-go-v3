package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolNodeAz 资源池中节点的AZ信息。
type PoolNodeAz struct {

	// **参数解释**：可用区名称。 **取值范围**：不涉及。
	Az string `json:"az"`

	// **参数解释**：可用区资源实例的数量。 **取值范围**：不涉及。
	Count int32 `json:"count"`
}

func (o PoolNodeAz) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolNodeAz struct{}"
	}

	return strings.Join([]string{"PoolNodeAz", string(data)}, " ")
}
