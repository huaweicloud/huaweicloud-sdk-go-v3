package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteQueueUserListRequest Request Object
type DeleteQueueUserListRequest struct {

	// **参数解释**： 集群ID。获取方式方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 必须是有效的dws集群ID。 **取值范围**： 36位UUID。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 队列名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	QueueName string `json:"queue_name"`

	Body *WorkloadQueueUserReq `json:"body,omitempty"`
}

func (o DeleteQueueUserListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteQueueUserListRequest struct{}"
	}

	return strings.Join([]string{"DeleteQueueUserListRequest", string(data)}, " ")
}
