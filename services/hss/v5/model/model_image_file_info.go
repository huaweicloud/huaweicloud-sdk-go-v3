package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageFileInfo 查询镜像无归属文件列表，文件信息
type ImageFileInfo struct {

	// 文件名称
	FileName *string `json:"file_name,omitempty"`

	// 文件路径
	FilePath *string `json:"file_path,omitempty"`

	// 文件大小
	Size *int32 `json:"size,omitempty"`
}

func (o ImageFileInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageFileInfo struct{}"
	}

	return strings.Join([]string{"ImageFileInfo", string(data)}, " ")
}
