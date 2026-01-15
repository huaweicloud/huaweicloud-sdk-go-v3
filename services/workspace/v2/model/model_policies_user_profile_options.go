package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesUserProfileOptions 配置文件漫游技术配置。
type PoliciesUserProfileOptions struct {

	// 配置项内容。json格式
	UserProfileRule *string `json:"user_profile_rule,omitempty"`

	// 排除常用文件夹。json格式
	RedirExcludeCommonFolders *string `json:"redir_exclude_common_folders,omitempty"`

	// 禁止自定义文件夹/copy base = 0,copy back = 0。json格式
	RedirExcludeCopy0s *string `json:"redir_exclude_copy0s,omitempty"`

	// 禁止自定义文件夹/copy base = 0,copy back = 1。json格式
	RedirExcludeCopy1s *string `json:"redir_exclude_copy1s,omitempty"`

	// 禁止自定义文件夹/copy base = 1,copy back = 0。json格式
	RedirExcludeCopy2s *string `json:"redir_exclude_copy2s,omitempty"`

	// 禁止自定义文件夹/copy base = 1,copy back = 1。json格式
	RedirExcludeCopy3s *string `json:"redir_exclude_copy3s,omitempty"`

	// 允许自定义文件夹 json格式
	RedirExcludeIncludes *string `json:"redir_exclude_includes,omitempty"`
}

func (o PoliciesUserProfileOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesUserProfileOptions struct{}"
	}

	return strings.Join([]string{"PoliciesUserProfileOptions", string(data)}, " ")
}
