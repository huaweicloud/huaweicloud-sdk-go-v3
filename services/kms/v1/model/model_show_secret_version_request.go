package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowSecretVersionRequest struct {
	// 凭据的资源标识符。

	SecretId string `json:"secret_id"`
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
