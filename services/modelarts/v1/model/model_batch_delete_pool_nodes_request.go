package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeletePoolNodesRequest Request Object
type BatchDeletePoolNodesRequest struct {

	// **参数解释**： 资源池名称。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	PoolName string `json:"pool_name"`

	Body *NodesDeletionRequest `json:"body,omitempty"`
}

func (o BatchDeletePoolNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeletePoolNodesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeletePoolNodesRequest", string(data)}, " ")
}
