package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePrivateSnatOption 创建SNAT规则的请求体。
type CreatePrivateSnatOption struct {

	// 私网NAT网关实例的ID。
	GatewayId string `json:"gateway_id"`

	// 功能说明：规则匹配的CIDR。取值约束：与virsubnet_id参数二选一。
	Cidr *string `json:"cidr,omitempty"`

	// 功能说明：规则匹配的子网的ID。 取值约束：与cidr参数二选一。
	VirsubnetId *string `json:"virsubnet_id,omitempty"`

	// SNAT规则的描述。长度范围小于等于255个字符，不能包含“<”和“>”。
	Description *string `json:"description,omitempty"`

	// 功能说明：中转IP的ID的列表。 取值约束：中转IP的ID个数不能超过1个。
	TransitIpIds []string `json:"transit_ip_ids"`
}

func (o CreatePrivateSnatOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateSnatOption struct{}"
	}

	return strings.Join([]string{"CreatePrivateSnatOption", string(data)}, " ")
}
