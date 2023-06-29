package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceInfo 实例信息。
type InstanceInfo struct {

	// 实例id。
	InstanceId *int64 `json:"instance_id,omitempty"`

	// 应用名称。
	BusinessName *string `json:"business_name,omitempty"`

	// 应用id。
	BusinessId *int64 `json:"business_id,omitempty"`

	// 组件名称。
	AppName *string `json:"app_name,omitempty"`

	// 主机名称。
	HostName *string `json:"host_name,omitempty"`

	// 实例名称。
	InstanceName *string `json:"instance_name,omitempty"`

	// 主机ip地址。
	IpAddress *string `json:"ip_address,omitempty"`

	// 环境ID。
	EnvId *int64 `json:"env_id,omitempty"`

	// Javaagent版本。
	AgentVersion *string `json:"agent_version,omitempty"`

	// 最后心跳时间。
	LastHeartbeat *int64 `json:"last_heartbeat,omitempty"`

	// 注册时间。
	RegisterTime *int64 `json:"register_time,omitempty"`

	// 最后修改用户id。
	LastModifyUserId *string `json:"last_modify_user_id,omitempty"`

	// 实例状态。
	InstanceStatus *int32 `json:"instance_status,omitempty"`

	// 最后修改用户名称。
	LastModifyUserName *string `json:"last_modify_user_name,omitempty"`

	// 最后修改时间。
	LastModifyTime *int64 `json:"last_modify_time,omitempty"`
}

func (o InstanceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceInfo struct{}"
	}

	return strings.Join([]string{"InstanceInfo", string(data)}, " ")
}
