package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UploadSignatureFileRequest struct {

	// 描述
	FileDesc *string `json:"file_desc,omitempty"`

	Body *UploadSignatureFileRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o UploadSignatureFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadSignatureFileRequest struct{}"
	}

	return strings.Join([]string{"UploadSignatureFileRequest", string(data)}, " ")
}
