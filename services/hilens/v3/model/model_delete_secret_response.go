package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSecretResponse Response Object
type DeleteSecretResponse struct {
	Secret         *SecretId `json:"secret,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o DeleteSecretResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecretResponse struct{}"
	}

	return strings.Join([]string{"DeleteSecretResponse", string(data)}, " ")
}
