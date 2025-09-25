package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckCredentialRequest Request Object
type CheckCredentialRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *CheckCredentialRequestBody `json:"body,omitempty"`
}

func (o CheckCredentialRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckCredentialRequest struct{}"
	}

	return strings.Join([]string{"CheckCredentialRequest", string(data)}, " ")
}
