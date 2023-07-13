package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateCredentialRequestBody struct {
	Credential *CreateCredentialRequestBodyCredential `json:"credential,omitempty"`
}

func (o CreateCredentialRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCredentialRequestBody struct{}"
	}

	return strings.Join([]string{"CreateCredentialRequestBody", string(data)}, " ")
}
