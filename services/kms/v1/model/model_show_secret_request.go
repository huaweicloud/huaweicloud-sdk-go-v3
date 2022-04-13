package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowSecretRequest struct {
	// 凭据的资源标识符。

	SecretId string `json:"secret_id"`
}

func (o ShowSecretRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecretRequest struct{}"
	}

	return strings.Join([]string{"ShowSecretRequest", string(data)}, " ")
}
