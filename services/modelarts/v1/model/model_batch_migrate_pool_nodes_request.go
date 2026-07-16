package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchMigratePoolNodesRequest Request Object
type BatchMigratePoolNodesRequest struct {

	// **参数解释**：资源池名称。该字段取自资源池metadata.name字段的值。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	PoolName string `json:"pool_name"`

	Body *NodeBatchMigrationRequest `json:"body,omitempty"`
}

func (o BatchMigratePoolNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchMigratePoolNodesRequest struct{}"
	}

	return strings.Join([]string{"BatchMigratePoolNodesRequest", string(data)}, " ")
}
