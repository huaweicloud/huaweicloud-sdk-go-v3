package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SwitchoverReplicaSetRequestBody struct {

	// **参数解释** 组ID。 **约束限制** 不涉及。 **取值范围** - 对于集群实例，该参数为shard组ID或config组ID。可以调用“查询实例列表和详情”接口获取。 - 对于副本集实例，不传该参数。 **默认取值** 不涉及。
	GroupId *string `json:"group_id,omitempty"`
}

func (o SwitchoverReplicaSetRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchoverReplicaSetRequestBody struct{}"
	}

	return strings.Join([]string{"SwitchoverReplicaSetRequestBody", string(data)}, " ")
}
