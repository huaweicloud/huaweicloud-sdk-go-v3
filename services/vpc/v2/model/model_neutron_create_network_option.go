package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronCreateNetworkOption 创建network对象
type NeutronCreateNetworkOption struct {

	// 功能说明：网络名称 取值范围：0-255个字符 约束：不能为admin_external_net
	Name *string `json:"name,omitempty"`

	// 功能说明：资源的管理状态 取值范围：true、false 约束：只支持true
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 功能说明：是否支持跨租户共享 取值范围：true、false
	Shared *bool `json:"shared,omitempty"`

	// 功能说明：扩展属性：网络类型。管理员属性，普通租户不可见，允许 geneve类型租户执行操作。 取值范围：vxlan, geneve
	ProvidernetworkType *string `json:"provider:network_type,omitempty"`

	// 功能说明：端口安全使能标记 取值范围：true(启用)、false(禁用) 约束：如果不使能，则network下所有虚机的安全组和dhcp防欺骗不生效
	PortSecurityEnabled *bool `json:"port_security_enabled,omitempty"`
}

func (o NeutronCreateNetworkOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateNetworkOption struct{}"
	}

	return strings.Join([]string{"NeutronCreateNetworkOption", string(data)}, " ")
}
