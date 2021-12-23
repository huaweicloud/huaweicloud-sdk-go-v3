package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateSecretResponse struct {
	Secret         *Secret `json:"secret,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateSecretResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecretResponse struct{}"
	}

	return strings.Join([]string{"UpdateSecretResponse", string(data)}, " ")
}
