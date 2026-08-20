package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowQueueDetailsRequest Request Object
type ShowQueueDetailsRequest struct {

	// **参数解释**： 实例ID。获取方法如下：调用[查询所有实例列表](ListInstancesDetails.xml)接口，从响应体中获取实例ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：  Vhost名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Vhost string `json:"vhost"`

	// **参数解释**： Queue名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Queue string `json:"queue"`
}

func (o ShowQueueDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQueueDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowQueueDetailsRequest", string(data)}, " ")
}
