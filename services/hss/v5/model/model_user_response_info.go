package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserResponseInfo 账号信息
type UserResponseInfo struct {

	// **参数解释**: Agent ID **取值范围**: 字符长度1-64位
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释**： 主机ID **取值范围**： 字符长度1-64位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 服务器名称 **取值范围**: 字符长度1-256位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 服务器IP **取值范围**: 字符长度1-128位
	HostIp *string `json:"host_ip,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**: 是否允许登录 **约束限制**: 不涉及 **取值范围**: - true：是 - false：否  **默认取值**: 不涉及
	LoginPermission *bool `json:"login_permission,omitempty"`

	// **参数解释**： 管理员权限 **取值范围**： - true：是 - false：否
	RootPermission *bool `json:"root_permission,omitempty"`

	// **参数解释**： 用户组 **取值范围**： 字符长度1-128位
	UserGroupName *string `json:"user_group_name,omitempty"`

	// **参数解释**： 用户目录 **取值范围**： 字符长度1-256位
	UserHomeDir *string `json:"user_home_dir,omitempty"`

	// **参数解释**: 用户启动shell **取值范围**: 字符长度1-128位
	Shell *string `json:"shell,omitempty"`

	// **参数解释**: 最近扫描时间 **取值范围**: 最小值0，最大值9223372036854775807
	RecentScanTime *int64 `json:"recent_scan_time,omitempty"`

	// **参数解释**: 首次扫描时间 **取值范围**: 最小值0，最大值9223372036854775807
	FirstScanTime *int64 `json:"first_scan_time,omitempty"`

	// **参数解释**: 容器ID **取值范围**: 字符长度1-128位
	ContainerId *string `json:"container_id,omitempty"`

	// **参数解释**： 容器实例名称，只有容器类型的告警有 **取值范围**： 字符长度1-256位
	ContainerName *string `json:"container_name,omitempty"`
}

func (o UserResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserResponseInfo struct{}"
	}

	return strings.Join([]string{"UserResponseInfo", string(data)}, " ")
}
