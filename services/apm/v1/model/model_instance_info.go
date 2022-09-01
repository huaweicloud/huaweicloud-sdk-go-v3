package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 实例信息
type InstanceInfo struct {

	// 实例id
	InstanceId *int64 `json:"instance_id,omitempty" xml:"instance_id"`

	// 应用名称
	BusinessName *string `json:"business_name,omitempty" xml:"business_name"`

	// 应用id
	BusinessId *int64 `json:"business_id,omitempty" xml:"business_id"`

	// 组件名称
	AppName *string `json:"app_name,omitempty" xml:"app_name"`

	// 主机名称
	HostName *string `json:"host_name,omitempty" xml:"host_name"`

	// 实例名称
	InstanceName *string `json:"instance_name,omitempty" xml:"instance_name"`

	// 主机ip地址
	IpAddress *string `json:"ip_address,omitempty" xml:"ip_address"`

	// 环境ID
	EnvId *int64 `json:"env_id,omitempty" xml:"env_id"`

	// Javaagent版本
	AgentVersion *string `json:"agent_version,omitempty" xml:"agent_version"`

	// 最后心跳时间
	LastHeartbeat *int64 `json:"last_heartbeat,omitempty" xml:"last_heartbeat"`

	// 注册时间
	RegisterTime *int64 `json:"register_time,omitempty" xml:"register_time"`

	// 最后修改用户id
	LastModifyUserId *string `json:"last_modify_user_id,omitempty" xml:"last_modify_user_id"`

	// 实例状态
	InstanceStatus *int32 `json:"instance_status,omitempty" xml:"instance_status"`

	// 最后修改用户名称
	LastModifyUserName *string `json:"last_modify_user_name,omitempty" xml:"last_modify_user_name"`

	// 最后修改时间
	LastModifyTime *int64 `json:"last_modify_time,omitempty" xml:"last_modify_time"`
}

func (o InstanceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceInfo struct{}"
	}

	return strings.Join([]string{"InstanceInfo", string(data)}, " ")
}
