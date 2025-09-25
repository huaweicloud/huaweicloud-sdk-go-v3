package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserChangeHistoryResponseInfo **参数解释** 账号变动历史信息
type UserChangeHistoryResponseInfo struct {

	// **参数解释** agent标识 **取值范围** 长度1-128
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释** 变更类型 **取值范围** - ADD ：添加 - DELETE ：删除 - MODIFY ： 修改
	ChangeType *string `json:"change_type,omitempty"`

	// **参数解释** 主机ID **取值范围** 长度1-128
	HostId *string `json:"host_id,omitempty"`

	// **参数解释** 服务器名称 **取值范围** 长度1-128
	HostName *string `json:"host_name,omitempty"`

	// **参数解释** 服务器私有IP **取值范围** 长度1-128
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释** 是否有登录权限 **取值范围** true: 具有登录权限 false: 不具有登录权限
	LoginPermission *bool `json:"login_permission,omitempty"`

	// **参数解释** 是否有root权限 **取值范围** true: 具有root权限 false: 不具有root权限
	RootPermission *bool `json:"root_permission,omitempty"`

	// **参数解释** 用户组名称 **取值范围** 长度1-128
	UserGroupName *string `json:"user_group_name,omitempty"`

	// **参数解释** 用户目录 **取值范围** 长度1-128
	UserHomeDir *string `json:"user_home_dir,omitempty"`

	// **参数解释** 用户启动shell **取值范围** 长度1-128
	Shell *string `json:"shell,omitempty"`

	// **参数解释** 用户名称 **取值范围** 长度1-128
	UserName *string `json:"user_name,omitempty"`

	// **参数解释** 到期时间，采用时间戳，默认毫秒 **取值范围** 长度0-4070880000000
	ExpireTime *int64 `json:"expire_time,omitempty"`

	// **参数解释** 账号增加、修改、删除等操作的变更时间 **取值范围** 长度0-4070880000000
	RecentScanTime *int64 `json:"recent_scan_time,omitempty"`
}

func (o UserChangeHistoryResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserChangeHistoryResponseInfo struct{}"
	}

	return strings.Join([]string{"UserChangeHistoryResponseInfo", string(data)}, " ")
}
