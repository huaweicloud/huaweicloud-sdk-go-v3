package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateCredentialRequestBodyCredential struct {
	Description *string `json:"description,omitempty"`
}

func (o CreateCredentialRequestBodyCredential) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCredentialRequestBodyCredential struct{}"
	}

	return strings.Join([]string{"CreateCredentialRequestBodyCredential", string(data)}, " ")
}
