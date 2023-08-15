package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachPolicyRequest Request Object
type AttachPolicyRequest struct {

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
