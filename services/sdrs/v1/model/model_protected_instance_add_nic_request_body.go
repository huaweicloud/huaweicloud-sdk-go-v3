package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 保护实例添加网卡请求体
type ProtectedInstanceAddNicRequestBody struct {
	// 添加网卡的子网ID。该参数是子网的network_id，和neutron_network_id的值保持一致。

	SubnetId string `json:"subnet_id"`
	// 添加网卡的安全组信息。默认为Sys-default安全组。

	SecurityGroups *[]SecurityGroupsParams `json:"security_groups,omitempty"`
	// IP地址，若无该参数表示自动分配IP地址。

	IpAddress *string `json:"ip_address,omitempty"`
}

func (o ProtectedInstanceAddNicRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectedInstanceAddNicRequestBody struct{}"
	}

	return strings.Join([]string{"ProtectedInstanceAddNicRequestBody", string(data)}, " ")
}
