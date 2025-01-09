package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NatGateway struct {

	// 网关实例的ID
	Id *string `json:"id,omitempty"`

	// 项目的ID。
	TenantId *string `json:"tenant_id,omitempty"`

	// 公网NAT网关实例的名字，长度限制为64。
	Name *string `json:"name,omitempty"`

	// 公网NAT网关实例的描述，长度限制为255。
	Description *string `json:"description,omitempty"`

	// 公网NAT网关的规格。取值为：“1”：小型，SNAT最大连接数10000；“2”：中型，SNAT最大连接数50000；“3”：大型，SNAT最大连接数200000；“4”：超大型，SNAT最大连接数1000000
	Spec *string `json:"spec,omitempty"`

	// 公网NAT网关实例的状态。 枚举值： ACTIVE PENDING_CREATE PENDING_UPDATE PENDING_DELETE INACTIVE
	Status *string `json:"status,omitempty"`

	// 公网NAT网关实例的名字，长度限制为64。 解冻/冻结状态。取值范围： \"true\"：解冻 \"false\"：冻结
	AdminStateUp *string `json:"admin_state_up,omitempty"`

	// 公网NAT网关实例的创建时间，遵循UTC时间，格式是yyyy-mm-ddThh:mm:ssZ。
	CreatedAt *string `json:"created_at,omitempty"`

	// VPC的id。
	RouterId *string `json:"router_id,omitempty"`

	// 公网NAT网关下行口（DVR的下一跳）所属的network id。
	InternalNetworkId *string `json:"internal_network_id,omitempty"`

	// 企业项目ID。创建公网NAT网关实例时，关联的企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o NatGateway) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NatGateway struct{}"
	}

	return strings.Join([]string{"NatGateway", string(data)}, " ")
}
