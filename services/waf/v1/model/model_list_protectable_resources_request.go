package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProtectableResourcesRequest Request Object
type ListProtectableResourcesRequest struct {

	// **参数解释：** 负载均衡器所在VPC ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	VpcId *string `json:"vpc_id,omitempty"`

	// **参数解释：** 租户region **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Region *string `json:"region,omitempty"`

	ResourceType string `json:"resource_type"`
}

func (o ListProtectableResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProtectableResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListProtectableResourcesRequest", string(data)}, " ")
}
