package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteServiceSpecificCredentialV5Request Request Object
type DeleteServiceSpecificCredentialV5Request struct {

	// 服务专属凭证ID。
	CredentialId string `json:"credential_id"`

	// IAM用户ID。
	UserId string `json:"user_id"`
}

func (o DeleteServiceSpecificCredentialV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServiceSpecificCredentialV5Request struct{}"
	}

	return strings.Join([]string{"DeleteServiceSpecificCredentialV5Request", string(data)}, " ")
}
