package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServerHpsClusterCapacity struct {

	// **参数解释**：规格名称。 **约束限制**：长度1-65536个字符。 **默认取值**：不涉及。
	Flavor *string `json:"flavor,omitempty"`

	// **参数解释**：可用区ID。 **约束限制**：长度1-65536个字符。 **默认取值**：不涉及。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// **参数解释**：超节点集群ID。 **约束限制**：长度1-65536个字符。 **默认取值**：不涉及。
	HyperinstanceClusterId *string `json:"hyperinstance_cluster_id,omitempty"`

	// **参数解释**：超节点集群名称。 **约束限制**：长度1-65536个字符。 **默认取值**：不涉及。
	HyperinstanceClusterName *string `json:"hyperinstance_cluster_name,omitempty"`

	// **参数解释**：资源规格。 **约束限制**：长度1-65536个字符。 **默认取值**：不涉及。
	ResourceFlavor *string `json:"resource_flavor,omitempty"`

	// **参数解释**：售罄状态。 **约束限制**：布尔值（true/false）。 **默认取值**：不涉及。
	IsSoldOut *bool `json:"is_sold_out,omitempty"`
}

func (o ServerHpsClusterCapacity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerHpsClusterCapacity struct{}"
	}

	return strings.Join([]string{"ServerHpsClusterCapacity", string(data)}, " ")
}
