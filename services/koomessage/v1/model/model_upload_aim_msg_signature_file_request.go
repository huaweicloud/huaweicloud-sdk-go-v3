package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadAimMsgSignatureFileRequest Request Object
type UploadAimMsgSignatureFileRequest struct {

	// 文件描述。
	FileDesc *string `json:"file_desc,omitempty"`

	// 请求体参数类型，该字段必须设置为：multipart/form-data。
	ContentType string `json:"Content-Type"`

	Body *UploadAimMsgSignatureFileRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o UploadAimMsgSignatureFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadAimMsgSignatureFileRequest struct{}"
	}

	return strings.Join([]string{"UploadAimMsgSignatureFileRequest", string(data)}, " ")
}
