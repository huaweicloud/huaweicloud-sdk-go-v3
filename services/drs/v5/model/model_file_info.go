package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FileInfo 文件信息
type FileInfo struct {

	// 文件名
	FileName *string `json:"file_name,omitempty"`

	// 文件修改时间
	LastModified *string `json:"last_modified,omitempty"`
}

func (o FileInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileInfo struct{}"
	}

	return strings.Join([]string{"FileInfo", string(data)}, " ")
}
