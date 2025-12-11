package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AntiVirusPaidHostResponseInfo 付费病毒查杀服务器信息
type AntiVirusPaidHostResponseInfo struct {

	// **参数解释**： 服务器（主机）的唯一标识ID **取值范围**： 字符长度1-64位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 服务器名称 **取值范围**: 字符长度1-256位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 主机所属服务器组的唯一标识ID **取值范围**: 字符长度0-64位
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释**： 弹性公网IP地址 **取值范围**： 字符长度1-256位，支持IPv4或IPv6格式（IPv4长度7-15位，IPv6长度15-39位）
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**： 服务器私有IP **取值范围**： 字符长度1-128位
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**: 主机上安装的杀毒Agent的唯一标识ID，用于关联主机与杀毒服务 **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释**: 操作系统类型 **取值范围**: 包含如下2种   - Linux：Linux   - Windows：Windows
	OsType *string `json:"os_type,omitempty"`
}

func (o AntiVirusPaidHostResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AntiVirusPaidHostResponseInfo struct{}"
	}

	return strings.Join([]string{"AntiVirusPaidHostResponseInfo", string(data)}, " ")
}
