package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadSecretBlobRequest Request Object
type UploadSecretBlobRequest struct {
	Body *UploadSecretBlobRequestBody `json:"body,omitempty"`
}

func (o UploadSecretBlobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadSecretBlobRequest struct{}"
	}

	return strings.Join([]string{"UploadSecretBlobRequest", string(data)}, " ")
}
