package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NetworkNodeInfo struct {

	// 节点id
	Id string `json:"id"`

	// 节点
	Name string `json:"name"`

	// 状态
	Status string `json:"status"`

	// ip地址
	IpAddr string `json:"ip_addr"`

	// ipv6地址
	Ipv6Addr *string `json:"ipv6_addr,omitempty"`

	// 私有ip地址
	PrivateIpAddr *string `json:"private_ip_addr,omitempty"`

	// 私有ipv6地址
	PrivateIpv6Addr *string `json:"private_ipv6_addr,omitempty"`

	// 运行时版本
	RuntimeVersion string `json:"runtime_version"`

	// os版本
	OsVersion string `json:"os_version"`

	// 安全组
	SecurityGroup string `json:"security_group"`

	// 服务器id
	ServerId *string `json:"server_id,omitempty"`

	// 服务器类型
	ServerType *string `json:"server_type,omitempty"`
}

func (o NetworkNodeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkNodeInfo struct{}"
	}

	return strings.Join([]string{"NetworkNodeInfo", string(data)}, " ")
}
