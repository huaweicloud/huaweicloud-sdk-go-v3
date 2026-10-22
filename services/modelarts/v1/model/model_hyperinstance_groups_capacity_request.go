package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HyperinstanceGroupsCapacityRequest struct {

	// **参数解释**：hyperinstance-group ID列表。 **默认取值**：不涉及。 **取值范围**：1-5
	HyperinstanceGroupIds []string `json:"hyperinstance_group_ids"`

	// **参数解释**：可用区。 **默认取值**：不涉及。
	AvailabilityZone string `json:"availability_zone"`

	// **参数解释**：规格名称。 **默认取值**：不涉及。
	ResourceFlavor string `json:"resource_flavor"`
}

func (o HyperinstanceGroupsCapacityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HyperinstanceGroupsCapacityRequest struct{}"
	}

	return strings.Join([]string{"HyperinstanceGroupsCapacityRequest", string(data)}, " ")
}
