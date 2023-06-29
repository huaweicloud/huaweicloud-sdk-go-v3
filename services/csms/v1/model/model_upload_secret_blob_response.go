package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadSecretBlobResponse Response Object
type UploadSecretBlobResponse struct {
	Secret         *Secret `json:"secret,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UploadSecretBlobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadSecretBlobResponse struct{}"
	}

	return strings.Join([]string{"UploadSecretBlobResponse", string(data)}, " ")
}
