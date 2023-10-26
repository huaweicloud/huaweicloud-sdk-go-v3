package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FsFileCount 所有文件数目
type FsFileCount struct {

	// 目录数目
	Dir int64 `json:"dir"`

	// 普通文件数目
	Regular int64 `json:"regular"`

	// 管道文件数目
	Pipe int64 `json:"pipe"`

	// 字符设备数目
	Char int64 `json:"char"`

	// 块设备数目
	Block int64 `json:"block"`

	// 套接字数目
	Socket int64 `json:"socket"`

	// 符号链接数目
	Symlink int64 `json:"symlink"`
}

func (o FsFileCount) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FsFileCount struct{}"
	}

	return strings.Join([]string{"FsFileCount", string(data)}, " ")
}
