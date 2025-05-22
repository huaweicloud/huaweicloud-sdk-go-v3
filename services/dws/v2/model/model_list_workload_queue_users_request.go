package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkloadQueueUsersRequest Request Object
type ListWorkloadQueueUsersRequest struct {

	// **参数解释**： 集群ID。获取方式方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 必须是有效的dws集群ID。 **取值范围**： 36位UUID。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 队列名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	QueueName string `json:"queue_name"`

	// **参数解释**： 分页单页大小。 **约束限制**： 不涉及。 **取值范围**： 大于0。 **默认取值**： 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 分页查询，偏移量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListWorkloadQueueUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkloadQueueUsersRequest struct{}"
	}

	return strings.Join([]string{"ListWorkloadQueueUsersRequest", string(data)}, " ")
}
