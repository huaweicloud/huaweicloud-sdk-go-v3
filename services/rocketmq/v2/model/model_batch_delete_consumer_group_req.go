package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteConsumerGroupReq struct {

	// **参数解释**： 待删除的消费组列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	JobId *string `json:"job_id,omitempty"`
}

func (o BatchDeleteConsumerGroupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteConsumerGroupReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteConsumerGroupReq", string(data)}, " ")
}
