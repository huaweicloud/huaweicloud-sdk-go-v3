package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DescribeGroupsRespGroup 消费组信息。
type DescribeGroupsRespGroup struct {

	// **参数解释**： 消费组状态。 **取值范围**： - Dead：消费组内没有任何成员，且没有任何元数据。 - Empty：消费组内没有任何成员，存在元数据。 - PreparingRebalance：准备开启rebalance。 - CompletingRebalance：所有成员加入group。 - Stable：消费组内成员可正常消费。
	State *string `json:"state,omitempty"`

	// **参数解释**： 创建时间。 **取值范围**： 不涉及。
	CreatedAt *int64 `json:"createdAt,omitempty"`

	// **参数解释**： 消费组名称。 **取值范围**： 不涉及。
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释**： 协调器编号。 **取值范围**： 不涉及。
	CoordinatorId *int32 `json:"coordinator_id,omitempty"`

	// **参数解释**： 分区分配策略。 **取值范围**： 不涉及。
	AssignmentStrategy *string `json:"assignment_strategy,omitempty"`

	// **参数解释**： 消费组描述。 **取值范围**： 不涉及。
	GroupDesc *string `json:"group_desc,omitempty"`
}

func (o DescribeGroupsRespGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DescribeGroupsRespGroup struct{}"
	}

	return strings.Join([]string{"DescribeGroupsRespGroup", string(data)}, " ")
}
