package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RedisSchema **参数解释**： 待重分布的schema对象。 **取值范围**： 不涉及。
type RedisSchema struct {

	// **参数解释**： schema名称。 **取值范围**： 不涉及。
	SchemaName *string `json:"schema_name,omitempty"`

	// **参数解释**： 优先级序号。 **取值范围**： 大于0的整数
	RedisOrder *int32 `json:"redis_order,omitempty"`
}

func (o RedisSchema) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RedisSchema struct{}"
	}

	return strings.Join([]string{"RedisSchema", string(data)}, " ")
}
