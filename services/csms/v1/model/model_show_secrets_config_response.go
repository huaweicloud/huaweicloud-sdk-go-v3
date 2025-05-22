package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecretsConfigResponse Response Object
type ShowSecretsConfigResponse struct {

	// 凭据强度检测配置是否打开。
	CheckingSecretStrength *bool `json:"checking_secret_strength,omitempty"`
	HttpStatusCode         int   `json:"-"`
}

func (o ShowSecretsConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecretsConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowSecretsConfigResponse", string(data)}, " ")
}
