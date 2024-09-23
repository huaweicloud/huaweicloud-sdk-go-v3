package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachPolicyRequest Request Object
type DetachPolicyRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 策略的唯一标识符（ID）。
	PolicyId string `json:"policy_id"`

	Body *PolicyTachReqBody `json:"body,omitempty"`
}

func (o DetachPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachPolicyRequest struct{}"
	}

	return strings.Join([]string{"DetachPolicyRequest", string(data)}, " ")
}
