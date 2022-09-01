package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateSecretRequest struct {

	// 凭据名称。
	SecretName string `json:"secret_name" xml:"secret_name"`

	Body *UpdateSecretRequestBody `json:"body,omitempty" xml:"body"`
}

func (o UpdateSecretRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecretRequest struct{}"
	}

	return strings.Join([]string{"UpdateSecretRequest", string(data)}, " ")
}
