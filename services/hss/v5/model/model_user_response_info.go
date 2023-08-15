package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserResponseInfo 账号信息
type UserResponseInfo struct {

	// agent_id
	AgentId *string `json:"agent_id,omitempty"`

	// 服务器ID
	HostId *string `json:"host_id,omitempty"`

	// 服务器名称
	HostName *string `json:"host_name,omitempty"`

	// 服务器ip
	HostIp *string `json:"host_ip,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// 是否有登陆权限
	LoginPermission *bool `json:"login_permission,omitempty"`

	// 是否有root权限
	RootPermission *bool `json:"root_permission,omitempty"`

	// 用户组
	UserGroupName *string `json:"user_group_name,omitempty"`

	// 用户目录
	UserHomeDir *string `json:"user_home_dir,omitempty"`

	// 用户启动shell
	Shell *string `json:"shell,omitempty"`

	// 到期时间，采用时间戳，默认毫秒，
	ExpireTime *int64 `json:"expire_time,omitempty"`

	// 最近扫描时间
	RecentScanTime *int64 `json:"recent_scan_time,omitempty"`
}

func (o UserResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserResponseInfo struct{}"
	}

	return strings.Join([]string{"UserResponseInfo", string(data)}, " ")
}
