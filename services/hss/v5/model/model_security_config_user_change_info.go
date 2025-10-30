package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityConfigUserChangeInfo 主机账户的历史变动记录
type SecurityConfigUserChangeInfo struct {

	// **参数解释**： 主机账户历史变动类型 **取值范围**： - add：添加 - delete：删除 - modify：修改
	ChangeType *string `json:"change_type,omitempty"`

	// **参数解释**： 是否有登录权限 **取值范围**： - true：有登录权限 - false：无登录权限
	LoginPermission *bool `json:"login_permission,omitempty"`

	// **参数解释**： 是否有root权限 **取值范围**： - true：有root权限 - false：无root权限
	RootPermission *bool `json:"root_permission,omitempty"`

	// **参数解释**： 用户组 **取值范围**： 不涉及
	UserGroupName *string `json:"user_group_name,omitempty"`

	// **参数解释**： 用户目录 **取值范围**： 不涉及
	UserHomeDir *string `json:"user_home_dir,omitempty"`

	// **参数解释**： 用户启动shell **取值范围**： 不涉及
	Shell *string `json:"shell,omitempty"`

	// **参数解释**： 用户名 **取值范围**： 不涉及
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**： 最新扫描时间，采用时间戳，默认毫秒 **取值范围**： 不涉及
	ScanTime *int64 `json:"scan_time,omitempty"`
}

func (o SecurityConfigUserChangeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityConfigUserChangeInfo struct{}"
	}

	return strings.Join([]string{"SecurityConfigUserChangeInfo", string(data)}, " ")
}
