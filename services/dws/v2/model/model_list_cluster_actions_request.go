package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterActionsRequest Request Object
type ListClusterActionsRequest struct {

	// **参数解释**： 集群ID。获取方式方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 任务名称。当前仅部分任务在操作中阶段支持查看任务详情。 **约束限制**： 不涉及。 **取值范围**： GROWING、RESIZE_FAILURE、RESTORING、RESTORING_FAILED、SNAPSHOTTING、SNAPSHOTTING_FAILED、FINE_GRAINED_RESTORING、FINE_GRAINED_RESTORING_FAILED **默认取值**： 不涉及。
	ActionName string `json:"action_name"`
}

func (o ListClusterActionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterActionsRequest struct{}"
	}

	return strings.Join([]string{"ListClusterActionsRequest", string(data)}, " ")
}
