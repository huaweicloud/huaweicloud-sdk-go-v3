package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadSignatureFileResponse Response Object
type UploadSignatureFileResponse struct {

	// 文件ID
	FileId         *string `json:"file_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UploadSignatureFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadSignatureFileResponse struct{}"
	}

	return strings.Join([]string{"UploadSignatureFileResponse", string(data)}, " ")
}
