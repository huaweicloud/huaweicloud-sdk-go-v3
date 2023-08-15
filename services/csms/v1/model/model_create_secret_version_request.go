package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecretVersionRequest Request Object
type CreateSecretVersionRequest struct {

	// 凭据名称。
	SecretName string `json:"secret_name"`

	Body *CreateSecretVersionRequestBody `json:"body,omitempty"`
}

func (o CreateSecretVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecretVersionRequest struct{}"
	}

	return strings.Join([]string{"CreateSecretVersionRequest", string(data)}, " ")
}
