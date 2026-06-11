package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartMenuConfigurations “开始菜单”文件夹重定向配置
type StartMenuConfigurations struct {

	// 配置文件夹重定向状态： 0: 未选取 1: 已选取
	StartMenuStatus *int32 `json:"start_menu_status,omitempty"`

	// 配置文件夹重定向类型： 0: 远程 1: 本地
	StartMenuRedirectionType *int32 `json:"start_menu_redirection_type,omitempty"`

	// 文件夹重定向(v2)用户存储路径。
	StartMenuStoragePath *string `json:"start_menu_storage_path,omitempty"`

	// 目标文件夹位置。
	StartMenuRelativePath *string `json:"start_menu_relative_path,omitempty"`

	// 是否开启用户对该文件夹的独占控制权限： 0: 禁用 1: 开启
	StartMenuExclusiveRights *int32 `json:"start_menu_exclusive_rights,omitempty"`

	// 启用文件夹重定向策略时，是否将现有内容移动到新位置： 0: 否 1: 是
	StartMenuMoveContents *int32 `json:"start_menu_move_contents,omitempty"`

	// 禁用或删除策略时，是否将内容移回本地用户配置文件位置： 0: 否 1: 是
	StartMenuMoveContentOnPolicyRemoval *int32 `json:"start_menu_move_content_on_policy_removal,omitempty"`
}

func (o StartMenuConfigurations) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartMenuConfigurations struct{}"
	}

	return strings.Join([]string{"StartMenuConfigurations", string(data)}, " ")
}
