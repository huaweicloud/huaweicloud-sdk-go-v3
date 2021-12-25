package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateSecretResponse struct {
	Secret         *SecretDetailResp `json:"secret,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o CreateSecretResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecretResponse struct{}"
	}

	return strings.Join([]string{"CreateSecretResponse", string(data)}, " ")
}
