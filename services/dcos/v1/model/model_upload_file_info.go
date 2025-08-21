package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadFileInfo 上传文件信息
type UploadFileInfo struct {

	// 用户上传附件名称
	FileName *string `json:"file_name,omitempty"`

	// 用户上传附件唯一id
	FileId string `json:"file_id"`

	// 文件类型
	FileType *string `json:"file_type,omitempty"`

	// 用户上传附件大小
	Size *int32 `json:"size,omitempty"`
}

func (o UploadFileInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadFileInfo struct{}"
	}

	return strings.Join([]string{"UploadFileInfo", string(data)}, " ")
}
