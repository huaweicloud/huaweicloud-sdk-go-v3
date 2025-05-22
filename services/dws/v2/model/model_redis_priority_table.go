package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RedisPriorityTable **参数解释**： 待重分布表优先级请求信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type RedisPriorityTable struct {

	// **参数解释**： 模式名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SchemaName *string `json:"schema_name,omitempty"`

	// **参数解释**： 表ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Id *int64 `json:"id,omitempty"`

	// **参数解释**： 表名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	TableName *string `json:"table_name,omitempty"`

	// **参数解释**： 优先级。 **约束限制**： 不涉及。 **取值范围**： 1~1024 **默认取值**： 不涉及。
	RedisOrder *int32 `json:"redis_order,omitempty"`
}

func (o RedisPriorityTable) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RedisPriorityTable struct{}"
	}

	return strings.Join([]string{"RedisPriorityTable", string(data)}, " ")
}
