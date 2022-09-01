package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UploadSecretBlobRequest struct {
	Body *UploadSecretBlobRequestBody `json:"body,omitempty" xml:"body"`
}

func (o UploadSecretBlobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadSecretBlobRequest struct{}"
	}

	return strings.Join([]string{"UploadSecretBlobRequest", string(data)}, " ")
}
