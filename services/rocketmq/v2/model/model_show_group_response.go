package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGroupResponse Response Object
type ShowGroupResponse struct {

	// **参数解释**： 是否可以消费。 **约束限制**： 不涉及。 **取值范围**： - true：可以消费。 - false：不可以消费。 **默认取值**： 不涉及。
	Enabled *bool `json:"enabled,omitempty"`

	// **参数解释**： 是否广播。 **约束限制**： 不涉及。 **取值范围**： - true：开启广播消费。 - false：不开启广播消费。 **默认取值**： 不涉及。
	Broadcast *bool `json:"broadcast,omitempty"`

	// **参数解释**： 关联的代理列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Brokers *[]string `json:"brokers,omitempty"`

	// **参数解释**： 消费组名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 消费组描述。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	GroupDesc *string `json:"group_desc,omitempty"`

	// **参数解释**： 创建时间。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	CreatedAt *int64 `json:"created_at,omitempty"`

	// **参数解释**： 最大重试次数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	RetryMaxTime *int32 `json:"retry_max_time,omitempty"`

	// **参数解释**： 权限。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Permissions *[]string `json:"permissions,omitempty"`

	// **参数解释**： 是否顺序消费。 **约束限制**： 不涉及。 **取值范围**： - true：开启顺序消费。 - false：不开启顺序消费。 **默认取值**： 不涉及。
	ConsumeOrderly *bool `json:"consume_orderly,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGroupResponse struct{}"
	}

	return strings.Join([]string{"ShowGroupResponse", string(data)}, " ")
}
