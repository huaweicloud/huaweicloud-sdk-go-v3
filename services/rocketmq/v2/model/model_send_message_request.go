package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SendMessageRequest Request Object
type SendMessageRequest struct {

	// **参数解释**： 引擎。 **约束限制**： 不涉及。 **取值范围**： - rocketmq：RocketMQ消息引擎。 - reliability：RocketMQ消息引擎别称。 **默认取值**： 不涉及。
	Engine string `json:"engine"`

	// **参数解释**： 实例ID。获取方法如下：调用“查询所有实例列表”接口，从响应体中获取实例ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId string `json:"instance_id"`

	Body *SendMessageReq `json:"body,omitempty"`
}

func (o SendMessageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendMessageRequest struct{}"
	}

	return strings.Join([]string{"SendMessageRequest", string(data)}, " ")
}
