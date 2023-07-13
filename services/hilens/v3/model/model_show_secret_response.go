package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecretResponse Response Object
type ShowSecretResponse struct {
	Secret         *SecretDetail `json:"secret,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowSecretResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecretResponse struct{}"
	}

	return strings.Join([]string{"ShowSecretResponse", string(data)}, " ")
}
