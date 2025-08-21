package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadFileRequest Request Object
type UploadFileRequest struct {
	Body *UploadFileRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o UploadFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadFileRequest struct{}"
	}

	return strings.Join([]string{"UploadFileRequest", string(data)}, " ")
}
