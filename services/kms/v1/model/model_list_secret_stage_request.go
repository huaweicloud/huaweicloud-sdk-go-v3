package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListSecretStageRequest struct {
	// 凭据的资源标识符。

	SecretId string `json:"secret_id"`
	// 凭据版本状态的名称。

	StageName string `json:"stage_name"`
}

func (o ListSecretStageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecretStageRequest struct{}"
	}

	return strings.Join([]string{"ListSecretStageRequest", string(data)}, " ")
}
