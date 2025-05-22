package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecretsConfigResponse Response Object
type UpdateSecretsConfigResponse struct {

	// 凭据强度检测配置是否打开。
	CheckingSecretStrength *bool `json:"checking_secret_strength,omitempty"`
	HttpStatusCode         int   `json:"-"`
}

func (o UpdateSecretsConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecretsConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateSecretsConfigResponse", string(data)}, " ")
}
