package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PartitionSpecContainerNetwork struct {

	// **参数解释**： 子网ID **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	SubnetID *string `json:"subnetID,omitempty"`
}

func (o PartitionSpecContainerNetwork) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PartitionSpecContainerNetwork struct{}"
	}

	return strings.Join([]string{"PartitionSpecContainerNetwork", string(data)}, " ")
}
