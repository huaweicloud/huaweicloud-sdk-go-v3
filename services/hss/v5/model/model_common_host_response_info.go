package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CommonHostResponseInfo 服务器信息
type CommonHostResponseInfo struct {

	// **参数解释**: 服务器的唯一标识ID **取值范围**: 字符长度1-64位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 服务器名称 **取值范围**: 字符长度1-256位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 服务器弹性IP地址 **取值范围**: IPv4格式（长度7-15位）、IPv6格式（长度15-39位）
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**: 服务器私有IP **取值范围**: 字符长度1-128位
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**: 主机上安装的杀毒Agent的唯一标识ID，用于关联主机与杀毒服务 **取值范围**: 字符长度1-64位
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释**： 操作系统类型 **取值范围**： - Linux：Linux - Windows：Windows
	OsType *string `json:"os_type,omitempty"`

	// **参数解释**： 服务器状态 **取值范围**： - ACTIVE：运行中 - SHUTOFF：关机 - BUILDING：创建中 - ERROR：故障
	HostStatus *string `json:"host_status,omitempty"`

	// **参数解释**： Agent状态 **取值范围**： - installed：已安装 - not_installed：未安装 - online：在线 - offline：离线 - install_failed：安装失败 - installing：安装中
	AgentStatus *string `json:"agent_status,omitempty"`

	// **参数解释**: 操作系统名称 **取值范围**: 字符长度1-64位
	OsName *string `json:"os_name,omitempty"`

	// **参数解释**: 操作系统版本 **取值范围**: 字符长度1-64位
	OsVersion *string `json:"os_version,omitempty"`

	// **参数解释**: 资产重要性 **约束限制**: 不涉及 **取值范围**： - important：重要资产 - common：一般资产 - test：测试资产  **默认取值**: 不涉及
	AssetValue *string `json:"asset_value,omitempty"`

	// **参数解释**: 集群ID **取值范围**: 字符长度1-64位
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**: 集群名称 **取值范围**: 字符长度1-256位
	ClusterName *string `json:"cluster_name,omitempty"`

	// **参数解释**: 服务器组ID **取值范围**: 字符范围1-64位
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释**: 集群名称 **取值范围**: 字符长度1-256位
	GroupName *string `json:"group_name,omitempty"`
}

func (o CommonHostResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommonHostResponseInfo struct{}"
	}

	return strings.Join([]string{"CommonHostResponseInfo", string(data)}, " ")
}
