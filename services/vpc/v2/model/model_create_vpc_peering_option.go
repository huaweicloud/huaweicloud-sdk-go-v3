package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// peering对象
type CreateVpcPeeringOption struct {

	// 功能说明：对等连接名称 取值范围：支持1~64个字符
	Name string `json:"name"`

	// 功能说明：对等连接的描述 取值范围：0-255个字符，不能包含“<”和“>”。
	Description *string `json:"description,omitempty"`

	RequestVpcInfo *VpcInfo `json:"request_vpc_info"`

	AcceptVpcInfo *VpcInfo `json:"accept_vpc_info"`
}

func (o CreateVpcPeeringOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpcPeeringOption struct{}"
	}

	return strings.Join([]string{"CreateVpcPeeringOption", string(data)}, " ")
}
