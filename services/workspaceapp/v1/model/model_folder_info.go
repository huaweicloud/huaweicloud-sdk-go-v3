package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FolderInfo 文件夹信息。
type FolderInfo struct {

	// 文件夹名称。
	Prefix *string `json:"prefix,omitempty"`

	// 文件inode。
	InodeNo *int64 `json:"inode_no,omitempty"`
}

func (o FolderInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FolderInfo struct{}"
	}

	return strings.Join([]string{"FolderInfo", string(data)}, " ")
}
