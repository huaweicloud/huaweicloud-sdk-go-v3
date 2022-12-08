package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateSasTokenRequest struct {

	// IAM用户的token，无需特殊权限。
	XAuthToken string `json:"X-Auth-Token"`

	// 该字段填为“application/json;charset=utf8”。
	ContentType string `json:"Content-Type"`

	Body *CreateSasTokenRequestBody `json:"body,omitempty"`
}

func (o CreateSasTokenRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSasTokenRequest struct{}"
	}

	return strings.Join([]string{"CreateSasTokenRequest", string(data)}, " ")
}
