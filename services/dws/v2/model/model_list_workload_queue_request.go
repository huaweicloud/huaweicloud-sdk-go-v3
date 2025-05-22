package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkloadQueueRequest Request Object
type ListWorkloadQueueRequest struct {

	// **参数解释**： 集群ID。获取方式方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 必须是有效的dws集群ID。 **取值范围**： 36位UUID。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 逻辑集群名称。逻辑集群模式下该字段必填。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	LogicalClusterName *string `json:"logical_cluster_name,omitempty"`
}

func (o ListWorkloadQueueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkloadQueueRequest struct{}"
	}

	return strings.Join([]string{"ListWorkloadQueueRequest", string(data)}, " ")
}
