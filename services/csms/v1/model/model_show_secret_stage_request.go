package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowSecretStageRequest struct {

	// 凭据名称。
	SecretName string `json:"secret_name"`

	// 凭据版本状态的名称。
	StageName string `json:"stage_name"`
}

func (o ShowSecretStageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecretStageRequest struct{}"
	}

	return strings.Join([]string{"ShowSecretStageRequest", string(data)}, " ")
}
