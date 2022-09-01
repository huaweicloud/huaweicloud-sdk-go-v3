package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateSecretTagRequest struct {

	// 凭据ID
	SecretId string `json:"secret_id" xml:"secret_id"`

	Body *CreateSecretTagRequestBody `json:"body,omitempty" xml:"body"`
}

func (o CreateSecretTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecretTagRequest struct{}"
	}

	return strings.Join([]string{"CreateSecretTagRequest", string(data)}, " ")
}
