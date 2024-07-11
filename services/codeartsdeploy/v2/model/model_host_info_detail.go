package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HostInfoDetail 主机详情信息
type HostInfoDetail struct {

	// 主机id
	HostId *string `json:"host_id,omitempty"`

	// 主机IP
	Ip *string `json:"ip,omitempty"`

	// 主机操作系统
	Os *string `json:"os,omitempty"`

	// 端口号
	Port *int32 `json:"port,omitempty"`

	Authorization *HostAuthorizationBody `json:"authorization,omitempty"`

	Permission *PermissionHostDetailNew `json:"permission,omitempty"`

	// 主机集群id
	GroupId *string `json:"group_id,omitempty"`

	// 主机名
	HostName *string `json:"host_name,omitempty"`

	// 是否为代理机
	AsProxy *bool `json:"as_proxy,omitempty"`

	// 代理机id
	ProxyHostId *string `json:"proxy_host_id,omitempty"`

	// 主机所属人名称
	OwnerName *string `json:"owner_name,omitempty"`

	ProxyHost *HostInfoDetail `json:"proxy_host,omitempty"`

	// 连通性状态
	ConnectionStatus *string `json:"connection_status,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 上次连通时间
	LastestConnectionTime *string `json:"lastest_connection_time,omitempty"`

	// 连通性验证结果
	ConnectionResult *string `json:"connection_result,omitempty"`

	// 免费启用应用运维服务（AOM），提供指标监控、日志查询、告警功能（自动安装数据采集器 ICAgent，仅支持华为云linux主机）
	InstallIcagent *bool `json:"install_icagent,omitempty"`

	// 创建人昵称
	NickName *string `json:"nick_name,omitempty"`
}

func (o HostInfoDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostInfoDetail struct{}"
	}

	return strings.Join([]string{"HostInfoDetail", string(data)}, " ")
}
