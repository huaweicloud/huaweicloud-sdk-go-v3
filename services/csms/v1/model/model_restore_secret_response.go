package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RestoreSecretResponse struct {
	Secret         *Secret `json:"secret,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestoreSecretResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreSecretResponse struct{}"
	}

	return strings.Join([]string{"RestoreSecretResponse", string(data)}, " ")
}
