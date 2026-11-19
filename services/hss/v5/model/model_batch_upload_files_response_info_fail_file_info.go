package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUploadFilesResponseInfoFailFileInfo 文件信息
type BatchUploadFilesResponseInfoFailFileInfo struct {

	// **参数解释**: 文件名称 **取值范围**: 字符长度1-512位
	FileName *string `json:"file_name,omitempty"`

	// **参数解释**: 失败原因 **取值范围**: 字符长度1-512位
	FailReason *string `json:"fail_reason,omitempty"`
}

func (o BatchUploadFilesResponseInfoFailFileInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUploadFilesResponseInfoFailFileInfo struct{}"
	}

	return strings.Join([]string{"BatchUploadFilesResponseInfoFailFileInfo", string(data)}, " ")
}
