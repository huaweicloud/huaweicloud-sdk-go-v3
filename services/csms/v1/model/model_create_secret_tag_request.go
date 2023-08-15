package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecretTagRequest Request Object
type CreateSecretTagRequest struct {

	// 凭据ID
	SecretId string `json:"secret_id"`

	Body *CreateSecretTagRequestBody `json:"body,omitempty"`
}

func (o CreateSecretTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecretTagRequest struct{}"
	}

	return strings.Join([]string{"CreateSecretTagRequest", string(data)}, " ")
}
