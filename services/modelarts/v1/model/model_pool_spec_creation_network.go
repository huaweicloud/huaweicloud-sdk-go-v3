package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolSpecCreationNetwork **参数解释**：资源池网络参数。创建物理资源池时必选。 **约束限制**：不涉及。
type PoolSpecCreationNetwork struct {

	// **参数解释**：网络名称，即网络详情中的metadata.name字段的值。用户接口通过指定网络名称创建网络，系统会自动创建子网，用户无法创建子网。默认将创建在第一个子网下。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Name *string `json:"name,omitempty"`
}

func (o PoolSpecCreationNetwork) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolSpecCreationNetwork struct{}"
	}

	return strings.Join([]string{"PoolSpecCreationNetwork", string(data)}, " ")
}
