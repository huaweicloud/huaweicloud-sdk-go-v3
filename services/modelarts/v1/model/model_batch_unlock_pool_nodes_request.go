package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUnlockPoolNodesRequest Request Object
type BatchUnlockPoolNodesRequest struct {

	// **参数解释**：资源池ID。取值资源池详情的metadata.name字段。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	PoolName string `json:"pool_name"`

	Body *BatchLockPoolNodesRequestBody `json:"body,omitempty"`
}

func (o BatchUnlockPoolNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUnlockPoolNodesRequest struct{}"
	}

	return strings.Join([]string{"BatchUnlockPoolNodesRequest", string(data)}, " ")
}
