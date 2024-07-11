package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnvironmentHostInfo 主机详情信息
type EnvironmentHostInfo struct {

	// 主机id
	HostId *string `json:"host_id,omitempty"`

	// 主机ip，如：161.17.101.12
	Ip *string `json:"ip,omitempty"`

	// ssh端口，如：22
	Port *int32 `json:"port,omitempty"`

	Permission *EnvironmentHostPermission `json:"permission,omitempty"`

	// 主机集群id
	GroupId *string `json:"group_id,omitempty"`

	// 主机名
	HostName *string `json:"host_name,omitempty"`

	// 是否为代理机
	AsProxy *bool `json:"as_proxy,omitempty"`

	// 代理机id
	ProxyHostId *string `json:"proxy_host_id,omitempty"`

	// 代理机名称
	ProxyHostName *string `json:"proxy_host_name,omitempty"`

	// 主机所属人id
	OwnerId *string `json:"owner_id,omitempty"`

	// 主机所属人名称
	OwnerName *string `json:"owner_name,omitempty"`

	// 连通性状态
	ConnectionStatus *string `json:"connection_status,omitempty"`

	// 上次连通时间
	LastestConnectionTime *string `json:"lastest_connection_time,omitempty"`

	// 连通性验证结果
	ConnectionResult *string `json:"connection_result,omitempty"`

	// 创建人昵称
	NickName *string `json:"nick_name,omitempty"`
}

func (o EnvironmentHostInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnvironmentHostInfo struct{}"
	}

	return strings.Join([]string{"EnvironmentHostInfo", string(data)}, " ")
}
