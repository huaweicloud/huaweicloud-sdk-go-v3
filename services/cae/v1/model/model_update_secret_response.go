package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecretResponse Response Object
type UpdateSecretResponse struct {

	// API版本，固定值“v1”，该值不可修改。
	ApiVersion *string `json:"api_version,omitempty"`

	// 资源种类。
	Kind *string `json:"kind,omitempty"`

	Spec           *SecretDetail `json:"spec,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o UpdateSecretResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecretResponse struct{}"
	}

	return strings.Join([]string{"UpdateSecretResponse", string(data)}, " ")
}
