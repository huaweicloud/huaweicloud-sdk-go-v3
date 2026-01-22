package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConsumerGroupOfTopicResponse Response Object
type ListConsumerGroupOfTopicResponse struct {

	// **参数解释**： 消费组列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Groups         *[]string `json:"groups,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListConsumerGroupOfTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConsumerGroupOfTopicResponse struct{}"
	}

	return strings.Join([]string{"ListConsumerGroupOfTopicResponse", string(data)}, " ")
}
