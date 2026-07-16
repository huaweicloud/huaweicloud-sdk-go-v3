package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetTopologiesRequest Request Object
type GetTopologiesRequest struct {

	// **参数解释**：Lite Server实例ID。 **约束限制**：^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：Lite Server实例对应的资源ID。 **约束限制**：^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$。 **取值范围**：不涉及。 **默认取值**：不涉及。
	ResourceId *string `json:"resource_id,omitempty"`
}

func (o GetTopologiesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetTopologiesRequest struct{}"
	}

	return strings.Join([]string{"GetTopologiesRequest", string(data)}, " ")
}
