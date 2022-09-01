package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateSecretStageRequest struct {

	// 凭据的资源标识符。
	SecretId string `json:"secret_id" xml:"secret_id"`

	// 凭据版本状态的名称。
	StageName string `json:"stage_name" xml:"stage_name"`

	Body *UpdateSecretStageRequestBody `json:"body,omitempty" xml:"body"`
}

func (o UpdateSecretStageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecretStageRequest struct{}"
	}

	return strings.Join([]string{"UpdateSecretStageRequest", string(data)}, " ")
}
