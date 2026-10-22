package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServerHpsGroupCapacity struct {

	// **参数解释**：可用区。 **默认取值**：不涉及。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// **参数解释**：hyperinstance-group ID。 **默认取值**：不涉及。
	HyperinstanceGroupId *string `json:"hyperinstance_group_id,omitempty"`

	// **参数解释**：hyperinstance-group名称。 **默认取值**：不涉及。
	HyperinstanceGroupName *string `json:"hyperinstance_group_name,omitempty"`

	// **参数解释**：超节点集群ID。 **默认取值**：不涉及。
	HyperinstanceClusterId *string `json:"hyperinstance_cluster_id,omitempty"`

	// **参数解释**：规格名称。 **默认取值**：不涉及。
	ResourceFlavor *string `json:"resource_flavor,omitempty"`

	// **参数解释**：是否售罄。 **默认取值**：不涉及。
	IsSoldOut *bool `json:"is_sold_out,omitempty"`
}

func (o ServerHpsGroupCapacity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerHpsGroupCapacity struct{}"
	}

	return strings.Join([]string{"ServerHpsGroupCapacity", string(data)}, " ")
}
