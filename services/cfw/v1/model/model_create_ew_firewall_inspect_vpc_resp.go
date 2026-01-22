package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateEwFirewallInspectVpcResp struct {

	// **参数解释**： 引流VPC的ID **取值范围**： 不涉及
	VpcId *string `json:"vpc_id,omitempty"`

	// **参数解释**： 创建的引流VPC的子网ID列表 **取值范围**： 不涉及
	SubnetIds *[]string `json:"subnet_ids,omitempty"`
}

func (o CreateEwFirewallInspectVpcResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEwFirewallInspectVpcResp struct{}"
	}

	return strings.Join([]string{"CreateEwFirewallInspectVpcResp", string(data)}, " ")
}
