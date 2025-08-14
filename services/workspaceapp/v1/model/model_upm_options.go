package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpmOptions 配置文件漫游技术配置。
type UpmOptions struct {

	// 配置项内容。
	UserProfileRule *string `json:"user_profile_rule,omitempty"`

	// 排除常用文件夹。
	RedirExcludeCommonFolders *string `json:"redir_exclude_common_folders,omitempty"`

	// 禁止自定义文件夹/copy base = 0,copy back = 0。
	RedirExcludeCopy0s *string `json:"redir_exclude_copy0s,omitempty"`

	// 禁止自定义文件夹/copy base = 0,copy back = 1。
	RedirExcludeCopy1s *string `json:"redir_exclude_copy1s,omitempty"`

	// 禁止自定义文件夹/copy base = 1,copy back = 0。
	RedirExcludeCopy2s *string `json:"redir_exclude_copy2s,omitempty"`

	// 禁止自定义文件夹/copy base = 1,copy back = 1。
	RedirExcludeCopy3s *string `json:"redir_exclude_copy3s,omitempty"`

	// 允许自定义文件夹
	RedirExcludeIncludes *string `json:"redir_exclude_includes,omitempty"`
}

func (o UpmOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpmOptions struct{}"
	}

	return strings.Join([]string{"UpmOptions", string(data)}, " ")
}
