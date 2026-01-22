package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConsumerGroupOrBatchDeleteConsumerGroupResponse Response Object
type CreateConsumerGroupOrBatchDeleteConsumerGroupResponse struct {

	// **参数解释**： 删除消费组的任务ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释**： 创建成功的消费组名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name           *string `json:"name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateConsumerGroupOrBatchDeleteConsumerGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConsumerGroupOrBatchDeleteConsumerGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateConsumerGroupOrBatchDeleteConsumerGroupResponse", string(data)}, " ")
}
