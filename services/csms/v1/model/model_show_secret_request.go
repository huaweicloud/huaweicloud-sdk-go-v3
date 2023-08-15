package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecretRequest Request Object
type ShowSecretRequest struct {

	// 凭据的名称。
	SecretName string `json:"secret_name"`
}

func (o ShowSecretRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecretRequest struct{}"
	}

	return strings.Join([]string{"ShowSecretRequest", string(data)}, " ")
}
