package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserVpcRequest **参数解释**：用户VPC配置。 **约束限制**：不涉及。
type UserVpcRequest struct {

	// **参数解释**：虚拟私有网络（VPC） ID。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	VpcId string `json:"vpc_id"`

	// **参数解释**：子网ID。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	SubnetId string `json:"subnet_id"`

	// **参数解释**：安全组ID列表。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	SecurityGroupIds []string `json:"security_group_ids"`

	// **参数解释**：连接的CIDR地址列表。 **约束限制**：选填参数，适用场景：用户希望通过挂载的网卡，访问其他网段的地址。 **取值范围**：不涉及。 **默认取值**：不涉及。
	ConnectCidrs *[]string `json:"connect_cidrs,omitempty"`

	// **参数解释**：NAT ID。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	NatId *string `json:"nat_id,omitempty"`

	// **参数解释**：EIP ID。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	EipId *string `json:"eip_id,omitempty"`
}

func (o UserVpcRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserVpcRequest struct{}"
	}

	return strings.Join([]string{"UserVpcRequest", string(data)}, " ")
}
