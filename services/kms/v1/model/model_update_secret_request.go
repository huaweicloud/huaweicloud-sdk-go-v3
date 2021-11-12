package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateSecretRequest struct {
	// 凭据对象唯一资源标识符。

	SecretId string `json:"secret_id"`

	Body *UpdateSecretRequestBody `json:"body,omitempty"`
}

func (o UpdateSecretRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecretRequest struct{}"
	}

	return strings.Join([]string{"UpdateSecretRequest", string(data)}, " ")
}
