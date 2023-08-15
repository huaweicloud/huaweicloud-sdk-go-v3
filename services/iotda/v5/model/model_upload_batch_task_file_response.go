package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadBatchTaskFileResponse Response Object
type UploadBatchTaskFileResponse struct {

	// 上传的批量任务文件ID，由平台自动生成。
	FileId *string `json:"file_id,omitempty"`

	// 上传的批量任务文件名称。
	FileName *string `json:"file_name,omitempty"`

	// 在物联网平台上传文件的时间。格式：yyyyMMdd'T'HHmmss'Z'，如20151212T121212Z。
	UploadTime     *string `json:"upload_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UploadBatchTaskFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadBatchTaskFileResponse struct{}"
	}

	return strings.Join([]string{"UploadBatchTaskFileResponse", string(data)}, " ")
}
