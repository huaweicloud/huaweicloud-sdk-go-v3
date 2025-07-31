package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUploadFilesResponseInfoSuccessFileInfo 文件信息
type BatchUploadFilesResponseInfoSuccessFileInfo struct {

	// 文件名称
	FileName *string `json:"file_name,omitempty"`

	// 文件ID，服务对保存的文件生成的唯一ID
	FileId *string `json:"file_id,omitempty"`
}

func (o BatchUploadFilesResponseInfoSuccessFileInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUploadFilesResponseInfoSuccessFileInfo struct{}"
	}

	return strings.Join([]string{"BatchUploadFilesResponseInfoSuccessFileInfo", string(data)}, " ")
}
