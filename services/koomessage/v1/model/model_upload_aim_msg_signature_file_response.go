package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadAimMsgSignatureFileResponse Response Object
type UploadAimMsgSignatureFileResponse struct {

	// 文件ID。
	FileId         *string `json:"file_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UploadAimMsgSignatureFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadAimMsgSignatureFileResponse struct{}"
	}

	return strings.Join([]string{"UploadAimMsgSignatureFileResponse", string(data)}, " ")
}
