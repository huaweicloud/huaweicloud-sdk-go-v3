package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodeDistributionOption struct {

	// **参数解释**: 对应可用区增加的只读节点数量。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	Num int32 `json:"num"`

	// **参数解释**: 可用分区。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	AvailabilityZone string `json:"availability_zone"`

	// **参数解释**: 规格码。 **约束限制**: 不涉及。 **取值范围**: 非空。 **默认取值**: 不涉及。
	FlavorRef string `json:"flavor_ref"`

	// **参数解释**: 只读参数组ID **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	ConfigurationId string `json:"configuration_id"`
}

func (o NodeDistributionOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeDistributionOption struct{}"
	}

	return strings.Join([]string{"NodeDistributionOption", string(data)}, " ")
}
