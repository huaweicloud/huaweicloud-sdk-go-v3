package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建托管专线参数对象
type CreateHostedDirectConnect struct {

	// 托管物理专线的名字。
	Name *string `json:"name,omitempty"`

	// 托管专线的描述信息
	Description *string `json:"description,omitempty"`

	// 指定托管专线接入带宽,单位Mbps
	Bandwidth *int32 `json:"bandwidth,omitempty"`

	// hosted物理专线对应的hosting物理专线的ID
	HostingId string `json:"hosting_id"`

	// 指定托管(hosted)专线预分配的vlan
	Vlan int32 `json:"vlan"`

	// 为其他租户创建托管专线，指定对应的租户ID
	ResourceTenantId string `json:"resource_tenant_id"`

	// 物理专线对端所在的物理位置，省/市/街道或IDC名字。
	PeerLocation *string `json:"peer_location,omitempty"`
}

func (o CreateHostedDirectConnect) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHostedDirectConnect struct{}"
	}

	return strings.Join([]string{"CreateHostedDirectConnect", string(data)}, " ")
}
