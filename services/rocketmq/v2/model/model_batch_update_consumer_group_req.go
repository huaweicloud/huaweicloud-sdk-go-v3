package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUpdateConsumerGroupReq struct {

	// **参数解释**： 消费组列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Groups *[]BatchUpdateConsumerGroup `json:"groups,omitempty"`
}

func (o BatchUpdateConsumerGroupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateConsumerGroupReq struct{}"
	}

	return strings.Join([]string{"BatchUpdateConsumerGroupReq", string(data)}, " ")
}
