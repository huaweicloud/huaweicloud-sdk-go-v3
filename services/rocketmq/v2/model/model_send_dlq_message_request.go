package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SendDlqMessageRequest Request Object
type SendDlqMessageRequest struct {

	// **参数解释**： 消息引擎。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Engine string `json:"engine"`

	// **参数解释**： 实例ID。获取方法如下：调用“查询所有实例列表”接口，从响应体中获取实例ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId string `json:"instance_id"`

	Body *DeadletterResendReq `json:"body,omitempty"`
}

func (o SendDlqMessageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendDlqMessageRequest struct{}"
	}

	return strings.Join([]string{"SendDlqMessageRequest", string(data)}, " ")
}
