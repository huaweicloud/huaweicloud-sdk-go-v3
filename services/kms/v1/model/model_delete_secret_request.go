package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteSecretRequest struct {

	// 凭据的资源标识符。
	SecretId string `json:"secret_id"`
}

func (o DeleteSecretRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecretRequest struct{}"
	}

	return strings.Join([]string{"DeleteSecretRequest", string(data)}, " ")
}
