package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConsumerConnectionsResponse Response Object
type ShowConsumerConnectionsResponse struct {

	// **参数解释**： 消费组名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释**： 消费组是否在线。 **约束限制**： 不涉及。 **取值范围**： - true：在线。 - false：不在线。 **默认取值**： 不涉及。
	Online *bool `json:"online,omitempty"`

	// **参数解释**： 订阅关系是否一致。 **约束限制**： 不涉及。 **取值范围**： - true：一致。 - false：不一致。 **默认取值**： 不涉及。
	SubscriptionConsistency *bool `json:"subscription_consistency,omitempty"`

	// **参数解释**： 消费者总数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**： 下个分页的offset。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	NextOffset *int32 `json:"next_offset,omitempty"`

	// **参数解释**： 上个分页的offset。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PreviousOffset *int32 `json:"previous_offset,omitempty"`

	// **参数解释**： 消费者订阅详情列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Clients        *[]ClientData `json:"clients,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowConsumerConnectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConsumerConnectionsResponse struct{}"
	}

	return strings.Join([]string{"ShowConsumerConnectionsResponse", string(data)}, " ")
}
