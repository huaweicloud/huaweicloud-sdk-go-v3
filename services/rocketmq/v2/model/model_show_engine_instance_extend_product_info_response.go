package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEngineInstanceExtendProductInfoResponse Response Object
type ShowEngineInstanceExtendProductInfoResponse struct {

	// **参数解释**： 总数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Total float32 `json:"total,omitempty"`

	// **参数解释**： 下个分页的offset。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	NextOffset *int32 `json:"next_offset,omitempty"`

	// **参数解释**： 上个分页的offset。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PreviousOffset *int32 `json:"previous_offset,omitempty"`

	// **参数解释**： 消息引擎类型。 **约束限制**： 不涉及。 **取值范围**： - rocketmq：RocketMQ消息引擎。 - reliability：RocketMQ消息引擎别称。 **默认取值**： 不涉及。
	Engine *string `json:"engine,omitempty"`

	// **参数解释**： 消息引擎支持的版本。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Versions *[]string `json:"versions,omitempty"`

	// **参数解释**： 规格变更的产品信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Products       *[]RocketMqExtendProductInfoEntity `json:"products,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o ShowEngineInstanceExtendProductInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEngineInstanceExtendProductInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowEngineInstanceExtendProductInfoResponse", string(data)}, " ")
}
