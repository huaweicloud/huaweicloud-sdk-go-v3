package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceLtCredentialRequest Request Object
type CreateInstanceLtCredentialRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	Body *CreateLongTermCredentialRequestBody `json:"body,omitempty"`
}

func (o CreateInstanceLtCredentialRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceLtCredentialRequest struct{}"
	}

	return strings.Join([]string{"CreateInstanceLtCredentialRequest", string(data)}, " ")
}
