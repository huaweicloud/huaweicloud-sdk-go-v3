package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateEwFirewallResp struct {

	// 东西向防护id
	Id *string `json:"id,omitempty"`

	Er *Er `json:"er,omitempty"`

	InspertionVpc *CreateEwFirewallInspectVpcResp `json:"inspertion_vpc,omitempty"`
}

func (o CreateEwFirewallResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEwFirewallResp struct{}"
	}

	return strings.Join([]string{"CreateEwFirewallResp", string(data)}, " ")
}
