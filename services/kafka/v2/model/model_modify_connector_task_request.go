package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyConnectorTaskRequest Request Object
type ModifyConnectorTaskRequest struct {

	// **参数解释**： 实例ID。获取方法如下：登录Kafka控制台，在Kafka实例详情页面查找实例ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**： 任务ID。获取方法如下：登录Kafka控制台，在Smart Connect页面查找任务ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	TaskId string `json:"task_id"`

	Body *SmartConnectTaskEntity `json:"body,omitempty"`
}

func (o ModifyConnectorTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyConnectorTaskRequest struct{}"
	}

	return strings.Join([]string{"ModifyConnectorTaskRequest", string(data)}, " ")
}
