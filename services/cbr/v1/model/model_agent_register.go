package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type AgentRegister struct {

	// 客户端ID
	AgentId string `json:"agent_id"`

	// 客户端所在的主机名
	HostName string `json:"host_name"`

	// 客户端所在主机的IP
	HostIp string `json:"host_ip"`

	// 客户端所在主机的操作系统
	HostOs string `json:"host_os"`

	// 客户端所在主机的主机别名
	HostNickname *string `json:"host_nickname,omitempty"`

	// 客户端版本
	AgentVersion *string `json:"agent_version,omitempty"`

	// 客户端类型，分本地客户端和云上客户端(cloud/native)
	AgentType *string `json:"agent_type,omitempty"`
}

func (o AgentRegister) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentRegister struct{}"
	}

	return strings.Join([]string{"AgentRegister", string(data)}, " ")
}
