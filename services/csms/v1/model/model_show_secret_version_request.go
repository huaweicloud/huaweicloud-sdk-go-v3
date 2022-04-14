package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowSecretVersionRequest struct {
	// 凭据名称。

	SecretName string `json:"secret_name"`
	// 凭据的版本标识符。

	VersionId string `json:"version_id"`
}

func (o ShowSecretVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecretVersionRequest struct{}"
	}

	return strings.Join([]string{"ShowSecretVersionRequest", string(data)}, " ")
}
