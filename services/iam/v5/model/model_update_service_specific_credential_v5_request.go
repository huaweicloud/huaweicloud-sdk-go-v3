package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateServiceSpecificCredentialV5Request Request Object
type UpdateServiceSpecificCredentialV5Request struct {

	// IAM用户ID。
	UserId string `json:"user_id"`

	// 服务专属凭证ID。
	CredentialId string `json:"credential_id"`

	Body *UpdateServiceSpecificCredentialReq `json:"body,omitempty"`
}

func (o UpdateServiceSpecificCredentialV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServiceSpecificCredentialV5Request struct{}"
	}

	return strings.Join([]string{"UpdateServiceSpecificCredentialV5Request", string(data)}, " ")
}
