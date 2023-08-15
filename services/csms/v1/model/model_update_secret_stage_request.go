package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecretStageRequest Request Object
type UpdateSecretStageRequest struct {

	// 凭据名称。
	SecretName string `json:"secret_name"`

	// 凭据版本状态的名称。满足 '^[a-zA-Z0-9._-]{1,64}$'
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
