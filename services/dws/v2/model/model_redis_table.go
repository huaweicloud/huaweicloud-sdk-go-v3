package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RedisTable **参数解释**： 重分布表信息。 **取值范围**： 不涉及。
type RedisTable struct {

	// **参数解释**： 表名。 **取值范围**： 不涉及。
	TableName *string `json:"table_name,omitempty"`

	// **参数解释**： 表唯一id。 **取值范围**： 不涉及。
	Id *int32 `json:"id,omitempty"`

	// **参数解释**： schema名。 **取值范围**： 不涉及。
	SchemaName *string `json:"schema_name,omitempty"`

	// **参数解释**： 逻辑集群名。 **取值范围**： 不涉及。
	LogicalClusterName *string `json:"logical_cluster_name,omitempty"`

	// **参数解释**： 表大小。 **取值范围**： 不涉及。
	Size *int64 `json:"size,omitempty"`

	// **参数解释**： 重分布类型。 **取值范围**： i：重分布中； y：重分布完成； n：未开始。
	Status *string `json:"status,omitempty"`
}

func (o RedisTable) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RedisTable struct{}"
	}

	return strings.Join([]string{"RedisTable", string(data)}, " ")
}
