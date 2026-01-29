package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScheduledTasksRequest Request Object
type ListScheduledTasksRequest struct {

	// **参数解释**： 实例ID。获取方法如下：调用[查询所有实例列表](ListInstancesDetails.xml)接口，从响应体中获取实例ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**： 开启查询的定时任务编号。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Start *string `json:"start,omitempty"`

	// **参数解释**： 查询的定时任务个数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Limit *string `json:"limit,omitempty"`

	// **参数解释**： 查询定时任务的最小时间，格式为YYYYMMDDHHmmss。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	BeginTime *string `json:"begin_time,omitempty"`

	// **参数解释**： 查询定时任务的最大时间，格式为YYYYMMDDHHmmss。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	EndTime *string `json:"end_time,omitempty"`
}

func (o ListScheduledTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScheduledTasksRequest struct{}"
	}

	return strings.Join([]string{"ListScheduledTasksRequest", string(data)}, " ")
}
