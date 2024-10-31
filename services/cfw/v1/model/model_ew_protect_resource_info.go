package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EwProtectResourceInfo 东西向防护的资源信息，例如VPC、VGW等
type EwProtectResourceInfo struct {

	// 防护资源类型：0 VPC，1 VGW，2 VPN，3 PEERING
	ProtectedResourceType int32 `json:"protected_resource_type"`

	// 防护资源名称
	ProtectedResourceName string `json:"protected_resource_name"`

	// 防护资源id
	ProtectedResourceId string `json:"protected_resource_id"`

	// 防护资源nat网关名称，专业版防火墙支持NAT规则，此字段表示防护连接的NAT的名称。
	ProtectedResourceNatName *string `json:"protected_resource_nat_name,omitempty"`

	// 防护资源nat网关id，专业版防火墙支持NAT规则，此字段表示防护连接的NAT的id。
	ProtectedResourceNatId *string `json:"protected_resource_nat_id,omitempty"`

	// 防火墙支持跨账户防护，此处为防护资源租户id
	ProtectedResourceProjectId *string `json:"protected_resource_project_id,omitempty"`

	// 防护资源模式，为er
	ProtectedResourceMode *string `json:"protected_resource_mode,omitempty"`

	// 防护资源的防护状态，0表示已关联，1表示未关联。
	Status *int32 `json:"status,omitempty"`
}

func (o EwProtectResourceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EwProtectResourceInfo struct{}"
	}

	return strings.Join([]string{"EwProtectResourceInfo", string(data)}, " ")
}
