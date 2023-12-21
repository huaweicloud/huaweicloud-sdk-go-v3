package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FsFileCount 所有文件数目
type FsFileCount struct {

	// 目录数目
	Dir *int64 `json:"dir,omitempty"`

	// 普通文件数目
	Regular *int64 `json:"regular,omitempty"`

	// 管道文件数目
	Pipe *int64 `json:"pipe,omitempty"`

	// 字符设备数目
	Char *int64 `json:"char,omitempty"`

	// 块设备数目
	Block *int64 `json:"block,omitempty"`

	// 套接字数目
	Socket *int64 `json:"socket,omitempty"`

	// 符号链接数目
	Symlink *int64 `json:"symlink,omitempty"`
}

func (o FsFileCount) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FsFileCount struct{}"
	}

	return strings.Join([]string{"FsFileCount", string(data)}, " ")
}
