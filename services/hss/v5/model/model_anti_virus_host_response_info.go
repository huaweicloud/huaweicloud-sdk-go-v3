package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AntiVirusHostResponseInfo 进程白名单可选服务器信息
type AntiVirusHostResponseInfo struct {

	// 主机ID
	HostId *string `json:"host_id,omitempty"`

	// 服务器名称
	HostName *string `json:"host_name,omitempty"`

	// 弹性公网IP地址
	PublicIp *string `json:"public_ip,omitempty"`

	// 服务器私有IP
	PrivateIp *string `json:"private_ip,omitempty"`

	// Agent ID
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
