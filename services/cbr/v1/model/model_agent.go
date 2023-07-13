package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Agent
type Agent struct {

	// 客户端创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 客户端更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	// 客户端ID
	AgentId string `json:"agent_id"`

	// 客户端版本号
	AgentVersion *string `json:"agent_version,omitempty"`

	// 客户端类型
	AgentType *string `json:"agent_type,omitempty"`

	// 客户端所在的主机名
	HostName *string `json:"host_name,omitempty"`

	// 客户端所在的主机昵称
	HostNickname *string `json:"host_nickname,omitempty"`

	// 客户端所在主机的IP
	HostIp *string `json:"host_ip,omitempty"`

	// 客户端主机所在的操作系统
	HostOs *string `json:"host_os,omitempty"`

	// 客户端状态
	Status *string `json:"status,omitempty"`

	// 客户端上次激活时间
	LastActiveTime *sdktime.SdkTime `json:"last_active_time,omitempty"`

	// 客户端的备份路径
	Paths *[]Path `json:"paths,omitempty"`
}

func (o Agent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Agent struct{}"
	}

	return strings.Join([]string{"Agent", string(data)}, " ")
}
