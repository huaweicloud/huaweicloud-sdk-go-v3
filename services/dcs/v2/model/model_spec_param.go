package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SpecParam struct {

	// **参数解释**： 实例分片数。 **约束限制**： 不涉及。 **取值范围**： 1-128。 **默认取值**： 不涉及。
	ShardingCount *int32 `json:"sharding_count,omitempty"`

	// 副本数 **参数解释**： 实例副本数。 **约束限制**： 不涉及。 **取值范围**： 1-10。 **默认取值**： 不涉及。
	ReplicaCount *int32 `json:"replica_count,omitempty"`

	// 副本数 **参数解释**： 实例类型。 **约束限制**： 不涉及。 **取值范围**： 1.ha：主备类型 2.single：单机类型 3.ha_rw_split：读写分离类型 4.proxy：proxy集群类型 5.cluster：cluster集群类型 **默认取值**： 不涉及。
	CacheMode *string `json:"cache_mode,omitempty"`
}

func (o SpecParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SpecParam struct{}"
	}

	return strings.Join([]string{"SpecParam", string(data)}, " ")
}
