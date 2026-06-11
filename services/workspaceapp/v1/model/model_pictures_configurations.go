package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PicturesConfigurations “图片”文件夹重定向配置
type PicturesConfigurations struct {

	// 配置文件夹重定向状态： 0: 未选取 1: 已选取
	PicturesStatus *int32 `json:"pictures_status,omitempty"`

	// 配置文件夹重定向类型： 0: 远程 1: 本地
	PicturesRedirectionType *int32 `json:"pictures_redirection_type,omitempty"`

	// 文件夹重定向(v2)用户存储路径。
	PicturesStoragePath *string `json:"pictures_storage_path,omitempty"`

	// 目标文件夹位置。
	PicturesRelativePath *string `json:"pictures_relative_path,omitempty"`

	// 是否开启用户对该文件夹的独占控制权限： 0: 禁用 1: 开启
	PicturesExclusiveRights *int32 `json:"pictures_exclusive_rights,omitempty"`

	// 启用文件夹重定向策略时，是否将现有内容移动到新位置： 0: 否 1: 是
	PicturesMoveContents *int32 `json:"pictures_move_contents,omitempty"`

	// 禁用或删除策略时，是否将内容移回本地用户配置文件位置： 0: 否 1: 是
	PicturesMoveContentOnPolicyRemoval *int32 `json:"pictures_move_content_on_policy_removal,omitempty"`
}

func (o PicturesConfigurations) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PicturesConfigurations struct{}"
	}

	return strings.Join([]string{"PicturesConfigurations", string(data)}, " ")
}
