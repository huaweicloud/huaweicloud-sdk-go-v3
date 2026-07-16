package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeBatchMigrationRequest **参数解释**：批量迁移节点请求体。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
type NodeBatchMigrationRequest struct {

	// **参数解释**：待迁移的节点名称列表。 **约束限制**：不涉及。
	MigrateNodeNames []string `json:"migrateNodeNames"`

	// **参数解释**：迁移起始集群名称。 专属算力资源时该字段与源资源池名称相同，取自源资源池metadata.name字段的值； 轻量算力集群时该字段取自迁移节点metadata.labels[os.modelarts.node/cluster]字段的值。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	FromClusterName string `json:"fromClusterName"`

	// **参数解释**：迁移目标集群名称。 专属算力资源时该字段与源资源池名称相同，取自目标资源池metadata.name字段的值； 轻量算力集群时该字段取自目标资源池内节点metadata.labels[os.modelarts.node/cluster]字段的值。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	ToClusterName string `json:"toClusterName"`

	// **参数解释**：迁移目标资源池名称。该字段取自目标资源池metadata.name字段的值。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	ToPoolName *string `json:"toPoolName,omitempty"`

	ResourceSpec *MigrateResourceSpec `json:"resourceSpec,omitempty"`
}

func (o NodeBatchMigrationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeBatchMigrationRequest struct{}"
	}

	return strings.Join([]string{"NodeBatchMigrationRequest", string(data)}, " ")
}
