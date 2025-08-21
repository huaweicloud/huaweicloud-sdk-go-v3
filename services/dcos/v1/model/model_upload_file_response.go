package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadFileResponse Response Object
type UploadFileResponse struct {

	// 用户上传附件名称
	FileName *string `json:"file_name,omitempty"`

	// 用户上传附件唯一id
	FileId *string `json:"file_id,omitempty"`

	// 文件类型
	FileType *string `json:"file_type,omitempty"`

	// 用户上传附件大小
	Size *int32 `json:"size,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UploadFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadFileResponse struct{}"
	}

	return strings.Join([]string{"UploadFileResponse", string(data)}, " ")
}
