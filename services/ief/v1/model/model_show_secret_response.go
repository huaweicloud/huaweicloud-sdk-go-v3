package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowSecretResponse struct {
	Secret         *SecretDetailResp `json:"secret,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowSecretResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecretResponse struct{}"
	}

	return strings.Join([]string{"ShowSecretResponse", string(data)}, " ")
}
