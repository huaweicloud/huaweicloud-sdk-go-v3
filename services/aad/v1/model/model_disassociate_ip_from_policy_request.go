package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateIpFromPolicyRequest Request Object
type DisassociateIpFromPolicyRequest struct {

	// 策略id
	PolicyId string `json:"policy_id"`

	Body *IpBindingBody `json:"body,omitempty"`
}

func (o DisassociateIpFromPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateIpFromPolicyRequest struct{}"
	}

	return strings.Join([]string{"DisassociateIpFromPolicyRequest", string(data)}, " ")
}
