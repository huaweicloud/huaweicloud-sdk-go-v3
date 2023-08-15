package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecretRequest Request Object
type UpdateSecretRequest struct {

	// 凭据名称。
	SecretName string `json:"secret_name"`

	Body *UpdateSecretRequestBody `json:"body,omitempty"`
}

func (o UpdateSecretRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecretRequest struct{}"
	}

	return strings.Join([]string{"UpdateSecretRequest", string(data)}, " ")
}
