package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstanceLtCredentialRequest Request Object
type DeleteInstanceLtCredentialRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 访问凭证ID，即token_id
	CredentialId string `json:"credential_id"`
}

func (o DeleteInstanceLtCredentialRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceLtCredentialRequest struct{}"
	}

	return strings.Join([]string{"DeleteInstanceLtCredentialRequest", string(data)}, " ")
}
