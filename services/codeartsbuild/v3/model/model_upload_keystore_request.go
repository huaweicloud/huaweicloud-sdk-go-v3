package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadKeystoreRequest Request Object
type UploadKeystoreRequest struct {
	Body *UploadKeystoreRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o UploadKeystoreRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadKeystoreRequest struct{}"
	}

	return strings.Join([]string{"UploadKeystoreRequest", string(data)}, " ")
}
