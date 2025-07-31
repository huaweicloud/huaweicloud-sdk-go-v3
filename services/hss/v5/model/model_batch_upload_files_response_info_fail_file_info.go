package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUploadFilesResponseInfoFailFileInfo 文件信息
type BatchUploadFilesResponseInfoFailFileInfo struct {

	// 文件名称
	FileName *string `json:"file_name,omitempty"`

	// 失败原因
	FailReason *string `json:"fail_reason,omitempty"`
}

func (o BatchUploadFilesResponseInfoFailFileInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUploadFilesResponseInfoFailFileInfo struct{}"
	}

	return strings.Join([]string{"BatchUploadFilesResponseInfoFailFileInfo", string(data)}, " ")
}
