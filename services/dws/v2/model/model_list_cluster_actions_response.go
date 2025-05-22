package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterActionsResponse Response Object
type ListClusterActionsResponse struct {

	// **参数解释**： 任务详情响应体。 **取值范围**： 随机生成的UUID。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 任务名称，同入参字段。 **取值范围**： 不涉及。
	ActionName *string `json:"action_name,omitempty"`

	// **参数解释**： 任务状态。 **取值范围**： 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 集群名称。 **取值范围**： 不涉及。
	ClusterName *string `json:"cluster_name,omitempty"`

	// **参数解释**： 任务提交时间。 **取值范围**： 不涉及。
	SubmitTime *string `json:"submit_time,omitempty"`

	// **参数解释**： 任务详情子项。 **取值范围**： 不涉及。
	Items          *[]ActionItemVo `json:"items,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListClusterActionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterActionsResponse struct{}"
	}

	return strings.Join([]string{"ListClusterActionsResponse", string(data)}, " ")
}
