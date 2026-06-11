package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LinksConfigurations “链接”文件夹重定向配置
type LinksConfigurations struct {

	// 配置文件夹重定向状态： 0: 未选取 1: 已选取
	LinksStatus *int32 `json:"links_status,omitempty"`

	// 配置文件夹重定向类型： 0: 远程 1: 本地
	LinksRedirectionType *int32 `json:"links_redirection_type,omitempty"`

	// 文件夹重定向(v2)用户存储路径。
	LinksStoragePath *string `json:"links_storage_path,omitempty"`

	// 目标文件夹位置。
	LinksRelativePath *string `json:"links_relative_path,omitempty"`

	// 是否开启用户对该文件夹的独占控制权限： 0: 禁用 1: 开启
	LinksExclusiveRights *int32 `json:"links_exclusive_rights,omitempty"`

	// 启用文件夹重定向策略时，是否将现有内容移动到新位置： 0: 否 1: 是
	LinksMoveContents *int32 `json:"links_move_contents,omitempty"`

	// 禁用或删除策略时，是否将内容移回本地用户配置文件位置： 0: 否 1: 是
	LinksMoveContentOnPolicyRemoval *int32 `json:"links_move_content_on_policy_removal,omitempty"`
}

func (o LinksConfigurations) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LinksConfigurations struct{}"
	}

	return strings.Join([]string{"LinksConfigurations", string(data)}, " ")
}
