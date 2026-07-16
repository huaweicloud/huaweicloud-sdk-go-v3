package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RoleReplica 角色副本配置
type RoleReplica struct {

	// **参数解释：** 角色名称。 **取值范围：** 不涉及。
	Name string `json:"name"`

	// **参数解释：** 最大副本数。 **取值范围：** 1~128。
	MaxReplicas *int32 `json:"max_replicas,omitempty"`

	// **参数解释：** 最小副本数。 **取值范围：** 1~128。
	MinReplicas *int32 `json:"min_replicas,omitempty"`
}

func (o RoleReplica) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RoleReplica struct{}"
	}

	return strings.Join([]string{"RoleReplica", string(data)}, " ")
}
