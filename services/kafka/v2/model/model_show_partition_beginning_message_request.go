package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPartitionBeginningMessageRequest Request Object
type ShowPartitionBeginningMessageRequest struct {

	// **参数解释**： 实例ID。获取方法如下：调用“查询所有实例列表”接口，从响应体中获取实例ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**： Topic名称。 **约束限制**： Topic名称必须以字母开头且只支持大小写字母、中横线、下划线以及数字。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Topic string `json:"topic"`

	// **参数解释**： 分区编号。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Partition int32 `json:"partition"`
}

func (o ShowPartitionBeginningMessageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPartitionBeginningMessageRequest struct{}"
	}

	return strings.Join([]string{"ShowPartitionBeginningMessageRequest", string(data)}, " ")
}
