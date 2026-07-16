package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolMetadataUpdateAnnotations **参数解释**：资源池的注释信息。
type PoolMetadataUpdateAnnotations struct {

	// **参数解释**：资源池描述信息，用于说明资源池用于某种指定场景。不能包含特殊字符!<>=&\"'。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	OsModelartsDescription *string `json:"os.modelarts/description,omitempty"`

	// **参数解释**：订单id，包周期创建和变更的时候需要传递该参数。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	OsModelartsOrderId *string `json:"os.modelarts/order.id,omitempty"`
}

func (o PoolMetadataUpdateAnnotations) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolMetadataUpdateAnnotations struct{}"
	}

	return strings.Join([]string{"PoolMetadataUpdateAnnotations", string(data)}, " ")
}
