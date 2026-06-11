package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VideosConfigurations “视频”文件夹重定向配置
type VideosConfigurations struct {

	// 配置文件夹重定向状态： 0: 未选取 1: 已选取
	VideosStatus *int32 `json:"videos_status,omitempty"`

	// 配置文件夹重定向类型： 0: 远程 1: 本地
	VideosRedirectionType *int32 `json:"videos_redirection_type,omitempty"`

	// 文件夹重定向(v2)用户存储路径。
	VideosStoragePath *string `json:"videos_storage_path,omitempty"`

	// 目标文件夹位置。
	VideosRelativePath *string `json:"videos_relative_path,omitempty"`

	// 是否开启用户对该文件夹的独占控制权限： 0: 禁用 1: 开启
	VideosExclusiveRights *int32 `json:"videos_exclusive_rights,omitempty"`

	// 启用文件夹重定向策略时，是否将现有内容移动到新位置： 0: 否 1: 是
	VideosMoveContents *int32 `json:"videos_move_contents,omitempty"`

	// 禁用或删除策略时，是否将内容移回本地用户配置文件位置： 0: 否 1: 是
	VideosMoveContentOnPolicyRemoval *int32 `json:"videos_move_content_on_policy_removal,omitempty"`
}

func (o VideosConfigurations) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideosConfigurations struct{}"
	}

	return strings.Join([]string{"VideosConfigurations", string(data)}, " ")
}
