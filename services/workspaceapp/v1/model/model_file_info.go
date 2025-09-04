package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FileInfo 文件信息。
type FileInfo struct {

	// 文件名称。
	FileName *string `json:"file_name,omitempty"`

	// 文件大小。
	Size *int64 `json:"size,omitempty"`

	// 文件的MIME类型。
	ContentType *string `json:"content_type,omitempty"`
}

func (o FileInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileInfo struct{}"
	}

	return strings.Join([]string{"FileInfo", string(data)}, " ")
}
