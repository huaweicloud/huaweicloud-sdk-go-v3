package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHttpPolicyRequest Request Object
type UpdateHttpPolicyRequest struct {

	// 防护策略id
	PolicyId string `json:"policy_id"`

	Body *UpdateHttpPolicyRequestBody `json:"body,omitempty"`
}

func (o UpdateHttpPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHttpPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateHttpPolicyRequest", string(data)}, " ")
}
