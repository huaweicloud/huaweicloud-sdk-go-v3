package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBackgroundTasksRequest Request Object
type ListBackgroundTasksRequest struct {

	// **参数解释**： 实例ID。获取方法如下：调用“查询所有实例列表”接口，从响应体中获取实例ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**： 开启查询的任务编号。 **约束限制**： 不涉及。 **取值范围**： 大于等于0。 **默认取值**： 不涉及。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 查询数量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 查询任务的最小时间，格式为YYYYMMDDHHmmss。 **约束限制**： 不涉及。 **取值范围**： 大于等于0。 **默认取值**： 不涉及。
	BeginTime *string `json:"begin_time,omitempty"`

	// **参数解释**： 查询任务的最大时间，格式为YYYYMMDDHHmmss。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	EndTime *string `json:"end_time,omitempty"`
}

func (o ListBackgroundTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackgroundTasksRequest struct{}"
	}

	return strings.Join([]string{"ListBackgroundTasksRequest", string(data)}, " ")
}
