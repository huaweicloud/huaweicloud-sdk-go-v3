package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateSecretVersionRequest struct {
	// 凭据的资源标识符。

	SecretId string `json:"secret_id"`

	Body *CreateSecretVersionRequestBody `json:"body,omitempty"`
}

func (o CreateSecretVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecretVersionRequest struct{}"
	}

	return strings.Join([]string{"CreateSecretVersionRequest", string(data)}, " ")
}
