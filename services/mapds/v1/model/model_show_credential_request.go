package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCredentialRequest Request Object
type ShowCredentialRequest struct {

	// IAM用户的token，无需特殊权限。
	XAuthToken string `json:"X-Auth-Token"`

	// 该字段填为“application/json;charset=utf8”。
	ContentType string `json:"Content-Type"`
}

func (o ShowCredentialRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCredentialRequest struct{}"
	}

	return strings.Join([]string{"ShowCredentialRequest", string(data)}, " ")
}
