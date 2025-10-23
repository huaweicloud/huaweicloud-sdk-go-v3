package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePolicyRequest Request Object
type UpdatePolicyRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	// 策略Id
	PolicyId string `json:"policy_id"`

	Body *CreateBccPolicyRequestBody `json:"body,omitempty"`
}

func (o UpdatePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdatePolicyRequest", string(data)}, " ")
}
