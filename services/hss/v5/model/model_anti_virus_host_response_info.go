package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AntiVirusHostResponseInfo 进程白名单可选服务器信息
type AntiVirusHostResponseInfo struct {

	// **参数解释**: 服务器ID **取值范围**: 字符长度1-64位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 服务器名称 **取值范围**: 字符长度1-256位
	HostName *string `json:"host_name,omitempty"`

	// 弹性公网IP地址
	PublicIp *string `json:"public_ip,omitempty"`

	// 服务器私有IP
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**: Agent ID **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	AgentId *string `json:"agent_id,omitempty"`

	// 操作系统类型，包含如下2种。   - Linux ：Linux   - Windows ：Winodws
	OsType *string `json:"os_type,omitempty"`

	// 服务器组ID
	GroupId *string `json:"group_id,omitempty"`
}

func (o AntiVirusHostResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AntiVirusHostResponseInfo struct{}"
	}

	return strings.Join([]string{"AntiVirusHostResponseInfo", string(data)}, " ")
}
