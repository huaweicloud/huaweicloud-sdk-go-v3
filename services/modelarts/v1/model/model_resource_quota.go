package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceQuota 资源配额参数。
type ResourceQuota struct {

	// **参数解释**：资源类型 **取值范围**：可选值如下： -  VPC：虚拟私有云。 -  SUBNET：子网。 -  SECURITY_GROUP：安全组。 -  SECURITY_GROUP_RULE：安全组规则。 -  PUBLIC_IP：公网IP。 -  VPC_PEER：VPC对端链接个数。 -  FIREWALL：防火墙。 -  SHARE_BANDWIDTH：共享带宽。 -  SHARE_BANDWIDTH_IP：共享带宽IP。 -  LOADBALANCER：负载均衡。 -  LISTENER：监听器。 -  PHYSICAL_CONNECT：物理连接。 -  VIRTUAL_INTERFACE：虚拟接口。 -  VPC_CONTAIN_ROUTETABLE：VPC包含的路由表。 -  ROUTETABLE_CONTAIN_ROUTES：路由表包含的路由。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 资源配额上限。 **取值范围**： 不涉及。
	Quota *string `json:"quota,omitempty"`

	// **参数解释**： 已使用配额。 **取值范围**： 不涉及。
	Used *string `json:"used,omitempty"`
}

func (o ResourceQuota) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceQuota struct{}"
	}

	return strings.Join([]string{"ResourceQuota", string(data)}, " ")
}
