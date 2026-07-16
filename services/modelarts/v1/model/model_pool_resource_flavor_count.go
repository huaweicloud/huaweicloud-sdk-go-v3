package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolResourceFlavorCount 资源池资源规格资源量数据模型。
type PoolResourceFlavorCount struct {

	// **参数解释**：资源规格ID。 **取值范围**：不涉及。
	Flavor string `json:"flavor"`

	// **参数解释**：资源池中资源规格实例数量。 **取值范围**：不涉及。
	Count int32 `json:"count"`

	// **参数解释**：资源池中资源规格实例弹性数量。物理池中该值和count相同。 **取值范围**：不涉及。
	MaxCount int32 `json:"maxCount"`

	// **参数解释**：资源池中期望创建的资源规格实例的az分布。
	Azs *[]PoolNodeAz `json:"azs,omitempty"`

	// **参数解释**：节点池ID。 **取值范围**：不涉及。
	NodePool *string `json:"nodePool,omitempty"`
}

func (o PoolResourceFlavorCount) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolResourceFlavorCount struct{}"
	}

	return strings.Join([]string{"PoolResourceFlavorCount", string(data)}, " ")
}
