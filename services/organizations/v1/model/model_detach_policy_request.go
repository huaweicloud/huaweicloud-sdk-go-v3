package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DetachPolicyRequest struct {

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
