package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RestoreSecretRequest struct {

	// 凭据名称。
	SecretName string `json:"secret_name" xml:"secret_name"`
}

func (o RestoreSecretRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreSecretRequest struct{}"
	}

	return strings.Join([]string{"RestoreSecretRequest", string(data)}, " ")
}
