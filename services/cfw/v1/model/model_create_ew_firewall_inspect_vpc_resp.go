package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateEwFirewallInspectVpcResp struct {

	// vpc id
	VpcId *string `json:"vpc_id,omitempty"`

	// 子网列表
	SubnetIds *[]string `json:"subnet_ids,omitempty"`
}

func (o CreateEwFirewallInspectVpcResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEwFirewallInspectVpcResp struct{}"
	}

	return strings.Join([]string{"CreateEwFirewallInspectVpcResp", string(data)}, " ")
}
