package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMessagesRequest Request Object
type ShowMessagesRequest struct {

	// **参数解释**： 实例ID。获取方法如下：登录Kafka控制台，在Kafka实例详情页面查找实例ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**： Topic名称。 **约束限制**： Topic名称必须以字母开头且只支持大小写字母、中横线、下划线以及数字。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Topic string `json:"topic"`

	// **参数解释**： 查询起始时间，为Unix时间戳格式。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 0。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**： 查询结束时间，为Unix时间戳格式。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 系统当前时间。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**： 单页返回消息数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 10。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 偏移量，表示从此偏移量开始查询。 **约束限制**： 不涉及。 **取值范围**： 大于等于0。 **默认取值**： 不涉及。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 分区编号。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 默认值为-1，若传入值为-1，则查询所有分区。
	Partition *string `json:"partition,omitempty"`
}

func (o ShowMessagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMessagesRequest struct{}"
	}

	return strings.Join([]string{"ShowMessagesRequest", string(data)}, " ")
}
