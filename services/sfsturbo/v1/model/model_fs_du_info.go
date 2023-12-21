package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FsDuInfo 目录资源使用情况(包含子目录)
type FsDuInfo struct {

	// 文件系统内合法的目录全路径
	Path *string `json:"path,omitempty"`

	// 占用容量，单位：byte
	UsedCapacity *int64 `json:"used_capacity,omitempty"`

	FileCount *FsFileCount `json:"file_count,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`
}

func (o FsDuInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FsDuInfo struct{}"
	}

	return strings.Join([]string{"FsDuInfo", string(data)}, " ")
}
