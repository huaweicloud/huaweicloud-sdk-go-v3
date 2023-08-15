package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSecretStageRequest Request Object
type DeleteSecretStageRequest struct {

	// 凭据的资源标识符。
	SecretName string `json:"secret_name"`

	// 凭据版本状态的名称。
	StageName string `json:"stage_name"`
}

func (o DeleteSecretStageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecretStageRequest struct{}"
	}

	return strings.Join([]string{"DeleteSecretStageRequest", string(data)}, " ")
}
