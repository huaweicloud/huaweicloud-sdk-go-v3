package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HostInfo 主机信息
type HostInfo struct {

	// 主机id
	Uuid *string `json:"uuid,omitempty"`

	// 主机IP
	Ip *string `json:"ip,omitempty"`

	// 主机操作系统
	Os *string `json:"os,omitempty"`

	// 端口
	Port *int32 `json:"port,omitempty"`

	Authorization *HostAuthorizationBody `json:"authorization,omitempty"`

	Permission *PermissionHostDetailNew `json:"permission,omitempty"`

	// 主机名称
	HostName *string `json:"host_name,omitempty"`

	// 是否为代理机
	AsProxy *bool `json:"as_proxy,omitempty"`

	// 主机集群id
	GroupId *string `json:"group_id,omitempty"`

	// 代理机id
	ProxyHostId *string `json:"proxy_host_id,omitempty"`

	// 主机所属人id
	OwnerId *string `json:"owner_id,omitempty"`

	// 主机所属人名称
	OwnerName *string `json:"owner_name,omitempty"`

	ProxyHost *HostInfo `json:"proxy_host,omitempty"`

	// 连通性状态
	ConnectionStatus *string `json:"connection_status,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 上次连通时间
	LastestConnectionTime *string `json:"lastest_connection_time,omitempty"`

	// 连通性验证结果
	ConnectionResult *string `json:"connection_result,omitempty"`

	// 主机所属人昵称
	NickName *string `json:"nick_name,omitempty"`

	// 导入状态
	ImportStatus *string `json:"import_status,omitempty"`

	// 关联环境数量
	EnvCount *int32 `json:"env_count,omitempty"`
}

func (o HostInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostInfo struct{}"
	}

	return strings.Join([]string{"HostInfo", string(data)}, " ")
}
