package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteConsumerGroupReq struct {

	// **参数解释**： 需要删除的消费组名称列表。 **约束限制**： 当批量删除消费组时必填。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Groups *[]string `json:"groups,omitempty"`
}

func (o BatchDeleteConsumerGroupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteConsumerGroupReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteConsumerGroupReq", string(data)}, " ")
}
