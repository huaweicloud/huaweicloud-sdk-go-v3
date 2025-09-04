package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FolderRedirectionOptions 文件夹重定向配置。
type FolderRedirectionOptions struct {

	// 目标文件夹位置。
	RelativePath *string `json:"relative_path,omitempty"`

	// 用户存储路径。
	StoragePaths *string `json:"storage_paths,omitempty"`

	// 文件夹选项。
	IncludeCommonFolders *int32 `json:"include_common_folders,omitempty"`
}

func (o FolderRedirectionOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FolderRedirectionOptions struct{}"
	}

	return strings.Join([]string{"FolderRedirectionOptions", string(data)}, " ")
}
