package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResidualResources struct {

	// **参数解释：** 负载均衡器监听器ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ElbListenerId *string `json:"elb_listener_id,omitempty"`

	// **参数解释：** 后端服务器组ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ElbPoolId *string `json:"elb_pool_id,omitempty"`

	// **参数解释：** 终端节点ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	VpcepId *string `json:"vpcep_id,omitempty"`
}

func (o ResidualResources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResidualResources struct{}"
	}

	return strings.Join([]string{"ResidualResources", string(data)}, " ")
}
