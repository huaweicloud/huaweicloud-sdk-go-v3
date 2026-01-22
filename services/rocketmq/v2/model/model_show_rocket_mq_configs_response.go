package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRocketMqConfigsResponse Response Object
type ShowRocketMqConfigsResponse struct {

	// **参数解释**： 总数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Total float32 `json:"total,omitempty"`

	// **参数解释**： 下个分页的offset。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	NextOffset *int32 `json:"next_offset,omitempty"`

	// **参数解释**： 上个分页的offset。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PreviousOffset *int32 `json:"previous_offset,omitempty"`

	// **参数解释**： RocketMQ配置。    **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	RocketmqConfigs *[]RocketMqConfigResp `json:"rocketmq_configs,omitempty"`
	HttpStatusCode  int                   `json:"-"`
}

func (o ShowRocketMqConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRocketMqConfigsResponse struct{}"
	}

	return strings.Join([]string{"ShowRocketMqConfigsResponse", string(data)}, " ")
}
