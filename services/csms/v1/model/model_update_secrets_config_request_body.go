package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateSecretsConfigRequestBody struct {

	// 凭据强度检测配置是否打开。
	CheckingSecretStrength *bool `json:"checking_secret_strength,omitempty"`
}

func (o UpdateSecretsConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecretsConfigRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSecretsConfigRequestBody", string(data)}, " ")
}
