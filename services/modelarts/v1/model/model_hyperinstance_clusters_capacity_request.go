package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HyperinstanceClustersCapacityRequest struct {

	// **参数解释**：超节点集群ID列表。 **约束限制**：数组长度0-5，每个元素长度1-128字符。 **默认取值**：不涉及。
	HyperinstanceClusterIds *[]string `json:"hyperinstance_cluster_ids,omitempty"`

	// **参数解释**：规格名称。 **约束限制**：长度1-65536个字符。 **默认取值**：不涉及。
	Flavor *string `json:"flavor,omitempty"`

	// **参数解释**：可用区。 **约束限制**：长度1-65536个字符。 **默认取值**：不涉及。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// **参数解释**：资源规格。 **约束限制**：长度1-65536个字符。 **默认取值**：不涉及。
	ResourceFlavor *string `json:"resource_flavor,omitempty"`
}

func (o HyperinstanceClustersCapacityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HyperinstanceClustersCapacityRequest struct{}"
	}

	return strings.Join([]string{"HyperinstanceClustersCapacityRequest", string(data)}, " ")
}
