package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ElbResp 集群ELB的相关信息
type ElbResp struct {

	// **参数解释**： 公网ip。 **取值范围**： 有效的公网ipv4地址。
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**： 内网ip。 **取值范围**： 有效的内网ipv4地址。
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**： Elb终端地址。 **取值范围**： 不涉及。
	PrivateEndpoint *string `json:"private_endpoint,omitempty"`

	// **参数解释**： Elb名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： Elb的ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： Elb所属VPC的ID。 **取值范围**： 不涉及。
	VpcId *string `json:"vpc_id,omitempty"`
}

func (o ElbResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ElbResp struct{}"
	}

	return strings.Join([]string{"ElbResp", string(data)}, " ")
}
