package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateSecretStageRequest struct {
	// 凭据名称。

	SecretName string `json:"secret_name"`
	// 凭据版本状态的名称。

	StageName string `json:"stage_name"`

	Body *UpdateSecretStageRequestBody `json:"body,omitempty"`
}

func (o UpdateSecretStageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecretStageRequest struct{}"
	}

	return strings.Join([]string{"UpdateSecretStageRequest", string(data)}, " ")
}
