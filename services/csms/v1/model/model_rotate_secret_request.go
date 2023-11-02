package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RotateSecretRequest Request Object
type RotateSecretRequest struct {

	// 凭据名称。
	SecretName string `json:"secret_name"`
}

func (o RotateSecretRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RotateSecretRequest struct{}"
	}

	return strings.Join([]string{"RotateSecretRequest", string(data)}, " ")
}
