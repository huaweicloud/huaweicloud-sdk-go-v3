package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConsumerGroupOrBatchDeleteConsumerGroupRequest Request Object
type CreateConsumerGroupOrBatchDeleteConsumerGroupRequest struct {

	// **参数解释**： 实例ID。获取方法如下：调用“查询所有实例列表”接口，从响应体中获取实例ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**： 批量删除消费组时使用，不配置则为创建接口。 **约束限制**： 不涉及。 **取值范围**： - delete：删除。 **默认取值**： 不涉及。
	Action *string `json:"action,omitempty"`

	Body *CreateConsumerGroupOrBatchDeleteConsumerGroupReq `json:"body,omitempty"`
}

func (o CreateConsumerGroupOrBatchDeleteConsumerGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConsumerGroupOrBatchDeleteConsumerGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateConsumerGroupOrBatchDeleteConsumerGroupRequest", string(data)}, " ")
}
