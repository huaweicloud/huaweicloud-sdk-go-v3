package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyConnectorTaskRequest Request Object
type ModifyConnectorTaskRequest struct {

	// **参数解释**： 实例ID。获取方法如下：调用“查询所有实例列表”接口，从响应体中获取实例ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**： 任务ID。获取方法如下：调用“查询Smart Connect任务列表”接口，从响应体中获取Smart Connect任务ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
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
