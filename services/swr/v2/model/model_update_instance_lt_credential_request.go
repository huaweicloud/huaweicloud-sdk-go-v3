package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceLtCredentialRequest Request Object
type UpdateInstanceLtCredentialRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 访问凭证ID，即token_id
	CredentialId string `json:"credential_id"`

	Body *UpdateLongTermCredentialRequestBody `json:"body,omitempty"`
}

func (o UpdateInstanceLtCredentialRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceLtCredentialRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceLtCredentialRequest", string(data)}, " ")
}
