package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConsumerGroup struct {

	// **参数解释**： 是否可以消费。 **取值范围**： - true：可以消费。 - false：不可以消费。
	Enabled *bool `json:"enabled,omitempty"`

	// **参数解释**： 是否广播。 **取值范围**： - true：可以广播。 - false：不可以广播。
	Broadcast *bool `json:"broadcast,omitempty"`

	// **参数解释**： 关联的代理列表。
	Brokers *[]string `json:"brokers,omitempty"`

	// **参数解释**： 消费组名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 消费组描述。 **取值范围**： 不涉及。
	GroupDesc *string `json:"group_desc,omitempty"`

	// **参数解释**： 最大重试次数。 **取值范围**： 1~16。
	RetryMaxTime *int32 `json:"retry_max_time,omitempty"`

	// **参数解释**： 创建时间戳。 **取值范围**： 不涉及。
	CreatedAt *int64 `json:"createdAt,omitempty"`

	// **参数解释**： 权限集。
	Permissions *[]string `json:"permissions,omitempty"`

	// **参数解释**： 是否按顺序消费。 **取值范围**： - true：按顺序消费。 - false：不按顺序消费。
	ConsumeOrderly *bool `json:"consume_orderly,omitempty"`
}

func (o ConsumerGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConsumerGroup struct{}"
	}

	return strings.Join([]string{"ConsumerGroup", string(data)}, " ")
}
