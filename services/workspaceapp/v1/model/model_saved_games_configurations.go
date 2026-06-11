package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SavedGamesConfigurations “保存的游戏”文件夹重定向配置
type SavedGamesConfigurations struct {

	// 配置文件夹重定向状态： 0: 未选取 1: 已选取
	SavedGamesStatus *int32 `json:"saved_games_status,omitempty"`

	// 配置文件夹重定向类型： 0: 远程 1: 本地
	SavedGamesRedirectionType *int32 `json:"saved_games_redirection_type,omitempty"`

	// 文件夹重定向(v2)用户存储路径。
	SavedGamesStoragePath *string `json:"saved_games_storage_path,omitempty"`

	// 目标文件夹位置。
	SavedGamesRelativePath *string `json:"saved_games_relative_path,omitempty"`

	// 是否开启用户对该文件夹的独占控制权限： 0: 禁用 1: 开启
	SavedGamesExclusiveRights *int32 `json:"saved_games_exclusive_rights,omitempty"`

	// 启用文件夹重定向策略时，是否将现有内容移动到新位置： 0: 否 1: 是
	SavedGamesMoveContents *int32 `json:"saved_games_move_contents,omitempty"`

	// 禁用或删除策略时，是否将内容移回本地用户配置文件位置： 0: 否 1: 是
	SavedGamesMoveContentOnPolicyRemoval *int32 `json:"saved_games_move_content_on_policy_removal,omitempty"`
}

func (o SavedGamesConfigurations) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SavedGamesConfigurations struct{}"
	}

	return strings.Join([]string{"SavedGamesConfigurations", string(data)}, " ")
}
