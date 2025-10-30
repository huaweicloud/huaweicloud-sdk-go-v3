package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageFileInfo 查询镜像无归属文件列表，文件信息
type ImageFileInfo struct {

	// **参数解释**: 文件名称 **取值范围**: 字符长度1-256
	FileName *string `json:"file_name,omitempty"`

	// **参数解释**: 文件路径 **取值范围**: 字符长度1-256
	FilePath *string `json:"file_path,omitempty"`

	// **参数解释**: 文件大小 **取值范围**: 最小值0，最大值65535
	Size *int32 `json:"size,omitempty"`
}

func (o ImageFileInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageFileInfo struct{}"
	}

	return strings.Join([]string{"ImageFileInfo", string(data)}, " ")
}
