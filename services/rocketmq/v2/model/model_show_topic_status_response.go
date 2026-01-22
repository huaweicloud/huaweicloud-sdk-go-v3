package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTopicStatusResponse Response Object
type ShowTopicStatusResponse struct {

	// **参数解释**： 最大偏移量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	MaxOffset *int32 `json:"max_offset,omitempty"`

	// **参数解释**： 最小偏移量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	MinOffset *int32 `json:"min_offset,omitempty"`

	// **参数解释**： 代理。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Brokers        *[]ShowTopicStatusRespBrokers `json:"brokers,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ShowTopicStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTopicStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowTopicStatusResponse", string(data)}, " ")
}
