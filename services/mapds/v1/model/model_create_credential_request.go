package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCredentialRequest Request Object
type CreateCredentialRequest struct {

	// IAM用户的token，无需特殊权限。
	XAuthToken string `json:"X-Auth-Token"`

	// 该字段填为“application/json;charset=utf8”。
	ContentType string `json:"Content-Type"`

	Body *CreateCredentialRequestBody `json:"body,omitempty"`
}

func (o CreateCredentialRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCredentialRequest struct{}"
	}

	return strings.Join([]string{"CreateCredentialRequest", string(data)}, " ")
}
