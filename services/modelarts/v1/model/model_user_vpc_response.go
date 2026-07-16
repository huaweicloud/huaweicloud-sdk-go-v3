package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserVpcResponse **参数解释**：用户VPC配置。
type UserVpcResponse struct {

	// **参数解释**：虚拟私有网络（VPC）ID。 **取值范围**：不涉及。
	VpcId *string `json:"vpc_id,omitempty"`

	// **参数解释**：子网ID。 **取值范围**：不涉及。
	SubnetId *string `json:"subnet_id,omitempty"`

	// **参数解释**：安全组ID列表。 **取值范围**：不涉及。
	SecurityGroupIds *[]string `json:"security_group_ids,omitempty"`

	// **参数解释**：连接的CIDR地址列表。 **取值范围**：不涉及。
	ConnectCidrs *string `json:"connect_cidrs,omitempty"`

	// **参数解释**：网卡ID。 **取值范围**：不涉及。
	PortId *[]string `json:"port_id,omitempty"`

	// **参数解释**：网卡ip。 **取值范围**：不涉及。
	PortIp *string `json:"port_ip,omitempty"`
}

func (o UserVpcResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserVpcResponse struct{}"
	}

	return strings.Join([]string{"UserVpcResponse", string(data)}, " ")
}
