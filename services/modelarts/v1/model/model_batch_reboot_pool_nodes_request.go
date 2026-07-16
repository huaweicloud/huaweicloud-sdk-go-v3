package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRebootPoolNodesRequest Request Object
type BatchRebootPoolNodesRequest struct {

	// **参数解释**：资源池名称。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	PoolName string `json:"pool_name"`

	Body *BatchRebootPoolNodesRequestBody `json:"body,omitempty"`
}

func (o BatchRebootPoolNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRebootPoolNodesRequest struct{}"
	}

	return strings.Join([]string{"BatchRebootPoolNodesRequest", string(data)}, " ")
}
