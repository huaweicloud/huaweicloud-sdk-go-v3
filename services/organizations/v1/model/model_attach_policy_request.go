package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachPolicyRequest Request Object
type AttachPolicyRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 策略的唯一标识符（ID）。
	PolicyId string `json:"policy_id"`

	Body *PolicyTachReqBody `json:"body,omitempty"`
}

func (o AttachPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachPolicyRequest struct{}"
	}

	return strings.Join([]string{"AttachPolicyRequest", string(data)}, " ")
}
